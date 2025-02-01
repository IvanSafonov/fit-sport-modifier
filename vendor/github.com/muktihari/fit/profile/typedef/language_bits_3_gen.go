// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LanguageBits3 uint8 // Base: uint8z

const (
	LanguageBits3Bulgarian LanguageBits3 = 0x01
	LanguageBits3Romanian  LanguageBits3 = 0x02
	LanguageBits3Chinese   LanguageBits3 = 0x04
	LanguageBits3Japanese  LanguageBits3 = 0x08
	LanguageBits3Korean    LanguageBits3 = 0x10
	LanguageBits3Taiwanese LanguageBits3 = 0x20
	LanguageBits3Thai      LanguageBits3 = 0x40
	LanguageBits3Hebrew    LanguageBits3 = 0x80
	LanguageBits3Invalid   LanguageBits3 = 0x0
)

func (l LanguageBits3) Uint8() uint8 { return uint8(l) }

func (l LanguageBits3) String() string {
	switch l {
	case LanguageBits3Bulgarian:
		return "bulgarian"
	case LanguageBits3Romanian:
		return "romanian"
	case LanguageBits3Chinese:
		return "chinese"
	case LanguageBits3Japanese:
		return "japanese"
	case LanguageBits3Korean:
		return "korean"
	case LanguageBits3Taiwanese:
		return "taiwanese"
	case LanguageBits3Thai:
		return "thai"
	case LanguageBits3Hebrew:
		return "hebrew"
	default:
		return "LanguageBits3Invalid(" + strconv.FormatUint(uint64(l), 10) + ")"
	}
}

// FromString parse string into LanguageBits3 constant it's represent, return LanguageBits3Invalid if not found.
func LanguageBits3FromString(s string) LanguageBits3 {
	switch s {
	case "bulgarian":
		return LanguageBits3Bulgarian
	case "romanian":
		return LanguageBits3Romanian
	case "chinese":
		return LanguageBits3Chinese
	case "japanese":
		return LanguageBits3Japanese
	case "korean":
		return LanguageBits3Korean
	case "taiwanese":
		return LanguageBits3Taiwanese
	case "thai":
		return LanguageBits3Thai
	case "hebrew":
		return LanguageBits3Hebrew
	default:
		return LanguageBits3Invalid
	}
}

// List returns all constants.
func ListLanguageBits3() []LanguageBits3 {
	return []LanguageBits3{
		LanguageBits3Bulgarian,
		LanguageBits3Romanian,
		LanguageBits3Chinese,
		LanguageBits3Japanese,
		LanguageBits3Korean,
		LanguageBits3Taiwanese,
		LanguageBits3Thai,
		LanguageBits3Hebrew,
	}
}
