package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/hash"
	"github.com/muktihari/fit/kit/hash/crc16"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func main() {
	args := Args{}
	if err := args.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	modifier := FitModifier{}
	if err := modifier.ProcessFile(args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type Args struct {
	SportName string
	Sport     int
	SubSport  int
	InName    string
	OutName   string
}

func (a *Args) Parse() error {
	a.Sport = -1
	a.SubSport = -1

	flag.StringVar(&a.SportName, "name", "", "new sport name")
	flag.IntVar(&a.Sport, "sport", -1, "new sport code")
	flag.IntVar(&a.SubSport, "subsport", -1, "new sub sport code")
	flag.Usage = Usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return errors.New("missing input fit file")
	}
	a.InName = args[0]

	if a.Sport > 255 {
		return errors.New("invalid sport option")
	}

	if a.SubSport > 255 {
		return errors.New("invalid subsport option")
	}

	if len(args) > 1 {
		a.OutName = args[1]
	}

	hasChangeOptions := a.SportName != "" || a.Sport > -1 || a.SubSport > -1
	if hasChangeOptions && a.OutName == "" {
		return errors.New("missing output fit file")
	}

	if !hasChangeOptions && a.OutName != "" {
		return errors.New("at least one option is required")
	}

	return nil
}

func Usage() {
	name := os.Args[0]
	fmt.Printf("Usage of %s [options] in.fit [out.fit]\n\n", name)
	fmt.Printf("Show current sport fields:\n%s in.fit\n\n", name)
	fmt.Printf("Replace sport name:\n%s -name \"XC Skate Ski\" ./in.fit ./out.fit\n\n", name)
	fmt.Printf("Replace sport name and sub sport code:\n%s -subsport 42 -name \"XC Skate Ski\" ./in.fit ./out.fit\n\n", name)
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

type FitModifier struct {
	args           Args
	fieldFactory   *factory.Factory
	msgDefinitions [proto.LocalMesgNumMask + 1]*proto.MessageDefinition
	crc            hash.Hash16
	out            *os.File
}

func (fm *FitModifier) ProcessFile(args Args) error {
	fm.args = args

	in, err := os.Open(args.InName)
	if err != nil {
		return fmt.Errorf("open fit: %w", err)
	}
	defer in.Close()

	if args.OutName != "" {
		fm.out, err = os.OpenFile(args.OutName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
		defer fm.out.Close()

		fm.crc = crc16.New(nil)
	}

	fm.fieldFactory = factory.New()

	dec := decoder.NewRaw()
	_, err = dec.Decode(in, fm.ProcessMessage)
	if err != nil {
		return err
	}

	if fm.out != nil {
		return fm.Write(binary.LittleEndian.AppendUint16(nil, fm.crc.Sum16()))
	}

	return nil
}

func (fm *FitModifier) ProcessMessage(flag decoder.RawFlag, b []byte) error {
	switch flag {
	case decoder.RawFlagFileHeader:
		break

	case decoder.RawFlagMesgDef:
		if err := fm.DecodeDefinition(&RawMsg{data: b}); err != nil {
			return err
		}

	case decoder.RawFlagMesgData:
		if err := fm.ProcessData(&RawMsg{data: b}); err != nil {
			return err
		}

	case decoder.RawFlagCRC:
		return nil
	}

	return fm.Write(b)
}

func (fm *FitModifier) Write(b []byte) error {
	if fm.out == nil {
		return nil
	}

	if _, err := fm.crc.Write(b); err != nil {
		return err
	}

	_, err := fm.out.Write(b)
	return err
}

func (fm *FitModifier) DecodeDefinition(rm *RawMsg) error {
	b, err := rm.Read(1)
	if err != nil {
		return err
	}

	header := b[0]
	b, err = rm.Read(5)
	if err != nil {
		return err
	}

	localMesgNum := header & proto.LocalMesgNumMask
	mesgDef := proto.MessageDefinition{}
	mesgDef.FieldDefinitions = make([]proto.FieldDefinition, 0, 255)
	mesgDef.Header = header
	mesgDef.Reserved = b[0]
	mesgDef.Architecture = b[1]
	if mesgDef.Architecture == proto.LittleEndian {
		mesgDef.MesgNum = typedef.MesgNum(binary.LittleEndian.Uint16(b[2:4]))
	} else {
		mesgDef.MesgNum = typedef.MesgNum(binary.BigEndian.Uint16(b[2:4]))
	}

	n := int(b[4])
	b, err = rm.Read(n * 3)
	if err != nil {
		return err
	}

	mesgDef.FieldDefinitions = mesgDef.FieldDefinitions[:0]
	var baseType basetype.BaseType
	for ; len(b) >= 3; b = b[3:] {
		baseType = basetype.BaseType(b[2])
		if !baseType.Valid() {
			return fmt.Errorf("message definition number: %s(%d): fields[%d].BaseType: %s: %w",
				mesgDef.MesgNum, mesgDef.MesgNum, len(mesgDef.FieldDefinitions), baseType, errors.New("errInvalidBaseType"))
		}
		mesgDef.FieldDefinitions = append(mesgDef.FieldDefinitions,
			proto.FieldDefinition{
				Num:      b[0],
				Size:     b[1],
				BaseType: baseType,
			})
	}

	fm.msgDefinitions[localMesgNum] = &mesgDef

	return nil
}

func (fm *FitModifier) ProcessData(rm *RawMsg) (err error) {
	b, err := rm.Read(1)
	if err != nil {
		return err
	}
	header := b[0]

	localMesgNum := header
	if (header & proto.MesgCompressedHeaderMask) == proto.MesgCompressedHeaderMask {
		localMesgNum = (header & proto.CompressedLocalMesgNumMask) >> proto.CompressedBitShift
	}
	mesgDef := fm.msgDefinitions[localMesgNum&proto.LocalMesgNumMask]
	if mesgDef == nil {
		return errors.New("ErrMesgDefMissing")
	}

	for i := range mesgDef.FieldDefinitions {
		fieldDef := &mesgDef.FieldDefinitions[i]
		field := fm.fieldFactory.CreateField(mesgDef.MesgNum, fieldDef.Num)

		if err = fm.ProcessField(mesgDef, fieldDef, field, rm); err != nil {
			return err
		}

		rm.Read(int(fieldDef.Size))
	}

	return nil
}

func (fm *FitModifier) ProcessField(mesgDef *proto.MessageDefinition, fieldDef *proto.FieldDefinition,
	field proto.Field, rm *RawMsg,
) (err error) {

	if mesgDef.MesgNum == typedef.MesgNumSport && field.Type == profile.SubSport {
		var data []byte
		if fm.args.SubSport > -1 {
			data = []byte{byte(fm.args.SubSport)}
		}

		current, err := rm.Peek(1)
		if err != nil {
			return err
		}
		return fm.ProcessFieldData("subsport", current[0], fm.args.SubSport, data, rm)
	}

	if mesgDef.MesgNum == typedef.MesgNumSport && field.Type == profile.Sport {
		var data []byte
		if fm.args.Sport > -1 {
			data = []byte{byte(fm.args.Sport)}
		}

		current, err := rm.Peek(1)
		if err != nil {
			return err
		}
		return fm.ProcessFieldData("sport", current[0], fm.args.Sport, data, rm)
	}

	if mesgDef.MesgNum == typedef.MesgNumSport && field.Name == "name" {
		var data []byte
		if fm.args.SportName != "" {
			data = make([]byte, int(fieldDef.Size))
			copy(data, fm.args.SportName)
		}

		current, err := rm.Peek(int(fieldDef.Size))
		if err != nil {
			return err
		}
		return fm.ProcessFieldData("name", string(current), fm.args.SportName, data, rm)
	}

	return nil
}

func (fm *FitModifier) ProcessFieldData(name string, current, updated any, updateData []byte, rm *RawMsg) error {
	if fm.out == nil || updateData == nil {
		fmt.Printf("%s: %v\n", name, current)
		return nil
	}

	fmt.Printf("%s: %v -> %v\n", name, current, updated)
	return rm.Replace(updateData)
}

type RawMsg struct {
	data []byte
	pos  int
}

func (r *RawMsg) Read(n int) ([]byte, error) {
	if n > len(r.data)-r.pos {
		return nil, fmt.Errorf("read %d bytes, %d available", n, len(r.data)-r.pos)
	}

	res := r.data[r.pos : n+r.pos]
	r.pos = n + r.pos
	return res, nil
}

func (r *RawMsg) Peek(n int) ([]byte, error) {
	if n > len(r.data)-r.pos {
		return nil, fmt.Errorf("peel %d bytes, %d available", n, len(r.data)-r.pos)
	}

	return r.data[r.pos : r.pos+n], nil
}

func (r *RawMsg) Replace(b []byte) error {
	if len(b) > len(r.data)-r.pos {
		return fmt.Errorf("replace %d bytes, %d available", len(b), len(r.data)-r.pos)
	}

	for i := 0; i < len(b); i++ {
		r.data[r.pos+i] = b[i]
	}

	return nil
}
