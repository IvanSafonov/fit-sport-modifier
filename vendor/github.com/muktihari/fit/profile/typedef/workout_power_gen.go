// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WorkoutPower uint32

const (
	WorkoutPowerWattsOffset WorkoutPower = 1000
	WorkoutPowerInvalid     WorkoutPower = 0xFFFFFFFF
)

func (w WorkoutPower) Uint32() uint32 { return uint32(w) }

func (w WorkoutPower) String() string {
	switch w {
	case WorkoutPowerWattsOffset:
		return "watts_offset"
	default:
		return "WorkoutPowerInvalid(" + strconv.FormatUint(uint64(w), 10) + ")"
	}
}

// FromString parse string into WorkoutPower constant it's represent, return WorkoutPowerInvalid if not found.
func WorkoutPowerFromString(s string) WorkoutPower {
	switch s {
	case "watts_offset":
		return WorkoutPowerWattsOffset
	default:
		return WorkoutPowerInvalid
	}
}

// List returns all constants.
func ListWorkoutPower() []WorkoutPower {
	return []WorkoutPower{
		WorkoutPowerWattsOffset,
	}
}
