// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type AntplusDeviceType uint8

const (
	AntplusDeviceTypeAntfs                   AntplusDeviceType = 1
	AntplusDeviceTypeBikePower               AntplusDeviceType = 11
	AntplusDeviceTypeEnvironmentSensorLegacy AntplusDeviceType = 12
	AntplusDeviceTypeMultiSportSpeedDistance AntplusDeviceType = 15
	AntplusDeviceTypeControl                 AntplusDeviceType = 16
	AntplusDeviceTypeFitnessEquipment        AntplusDeviceType = 17
	AntplusDeviceTypeBloodPressure           AntplusDeviceType = 18
	AntplusDeviceTypeGeocacheNode            AntplusDeviceType = 19
	AntplusDeviceTypeLightElectricVehicle    AntplusDeviceType = 20
	AntplusDeviceTypeEnvSensor               AntplusDeviceType = 25
	AntplusDeviceTypeRacquet                 AntplusDeviceType = 26
	AntplusDeviceTypeControlHub              AntplusDeviceType = 27
	AntplusDeviceTypeMuscleOxygen            AntplusDeviceType = 31
	AntplusDeviceTypeShifting                AntplusDeviceType = 34
	AntplusDeviceTypeBikeLightMain           AntplusDeviceType = 35
	AntplusDeviceTypeBikeLightShared         AntplusDeviceType = 36
	AntplusDeviceTypeExd                     AntplusDeviceType = 38
	AntplusDeviceTypeBikeRadar               AntplusDeviceType = 40
	AntplusDeviceTypeBikeAero                AntplusDeviceType = 46
	AntplusDeviceTypeWeightScale             AntplusDeviceType = 119
	AntplusDeviceTypeHeartRate               AntplusDeviceType = 120
	AntplusDeviceTypeBikeSpeedCadence        AntplusDeviceType = 121
	AntplusDeviceTypeBikeCadence             AntplusDeviceType = 122
	AntplusDeviceTypeBikeSpeed               AntplusDeviceType = 123
	AntplusDeviceTypeStrideSpeedDistance     AntplusDeviceType = 124
	AntplusDeviceTypeInvalid                 AntplusDeviceType = 0xFF
)

func (a AntplusDeviceType) Uint8() uint8 { return uint8(a) }

func (a AntplusDeviceType) String() string {
	switch a {
	case AntplusDeviceTypeAntfs:
		return "antfs"
	case AntplusDeviceTypeBikePower:
		return "bike_power"
	case AntplusDeviceTypeEnvironmentSensorLegacy:
		return "environment_sensor_legacy"
	case AntplusDeviceTypeMultiSportSpeedDistance:
		return "multi_sport_speed_distance"
	case AntplusDeviceTypeControl:
		return "control"
	case AntplusDeviceTypeFitnessEquipment:
		return "fitness_equipment"
	case AntplusDeviceTypeBloodPressure:
		return "blood_pressure"
	case AntplusDeviceTypeGeocacheNode:
		return "geocache_node"
	case AntplusDeviceTypeLightElectricVehicle:
		return "light_electric_vehicle"
	case AntplusDeviceTypeEnvSensor:
		return "env_sensor"
	case AntplusDeviceTypeRacquet:
		return "racquet"
	case AntplusDeviceTypeControlHub:
		return "control_hub"
	case AntplusDeviceTypeMuscleOxygen:
		return "muscle_oxygen"
	case AntplusDeviceTypeShifting:
		return "shifting"
	case AntplusDeviceTypeBikeLightMain:
		return "bike_light_main"
	case AntplusDeviceTypeBikeLightShared:
		return "bike_light_shared"
	case AntplusDeviceTypeExd:
		return "exd"
	case AntplusDeviceTypeBikeRadar:
		return "bike_radar"
	case AntplusDeviceTypeBikeAero:
		return "bike_aero"
	case AntplusDeviceTypeWeightScale:
		return "weight_scale"
	case AntplusDeviceTypeHeartRate:
		return "heart_rate"
	case AntplusDeviceTypeBikeSpeedCadence:
		return "bike_speed_cadence"
	case AntplusDeviceTypeBikeCadence:
		return "bike_cadence"
	case AntplusDeviceTypeBikeSpeed:
		return "bike_speed"
	case AntplusDeviceTypeStrideSpeedDistance:
		return "stride_speed_distance"
	default:
		return "AntplusDeviceTypeInvalid(" + strconv.FormatUint(uint64(a), 10) + ")"
	}
}

// FromString parse string into AntplusDeviceType constant it's represent, return AntplusDeviceTypeInvalid if not found.
func AntplusDeviceTypeFromString(s string) AntplusDeviceType {
	switch s {
	case "antfs":
		return AntplusDeviceTypeAntfs
	case "bike_power":
		return AntplusDeviceTypeBikePower
	case "environment_sensor_legacy":
		return AntplusDeviceTypeEnvironmentSensorLegacy
	case "multi_sport_speed_distance":
		return AntplusDeviceTypeMultiSportSpeedDistance
	case "control":
		return AntplusDeviceTypeControl
	case "fitness_equipment":
		return AntplusDeviceTypeFitnessEquipment
	case "blood_pressure":
		return AntplusDeviceTypeBloodPressure
	case "geocache_node":
		return AntplusDeviceTypeGeocacheNode
	case "light_electric_vehicle":
		return AntplusDeviceTypeLightElectricVehicle
	case "env_sensor":
		return AntplusDeviceTypeEnvSensor
	case "racquet":
		return AntplusDeviceTypeRacquet
	case "control_hub":
		return AntplusDeviceTypeControlHub
	case "muscle_oxygen":
		return AntplusDeviceTypeMuscleOxygen
	case "shifting":
		return AntplusDeviceTypeShifting
	case "bike_light_main":
		return AntplusDeviceTypeBikeLightMain
	case "bike_light_shared":
		return AntplusDeviceTypeBikeLightShared
	case "exd":
		return AntplusDeviceTypeExd
	case "bike_radar":
		return AntplusDeviceTypeBikeRadar
	case "bike_aero":
		return AntplusDeviceTypeBikeAero
	case "weight_scale":
		return AntplusDeviceTypeWeightScale
	case "heart_rate":
		return AntplusDeviceTypeHeartRate
	case "bike_speed_cadence":
		return AntplusDeviceTypeBikeSpeedCadence
	case "bike_cadence":
		return AntplusDeviceTypeBikeCadence
	case "bike_speed":
		return AntplusDeviceTypeBikeSpeed
	case "stride_speed_distance":
		return AntplusDeviceTypeStrideSpeedDistance
	default:
		return AntplusDeviceTypeInvalid
	}
}

// List returns all constants.
func ListAntplusDeviceType() []AntplusDeviceType {
	return []AntplusDeviceType{
		AntplusDeviceTypeAntfs,
		AntplusDeviceTypeBikePower,
		AntplusDeviceTypeEnvironmentSensorLegacy,
		AntplusDeviceTypeMultiSportSpeedDistance,
		AntplusDeviceTypeControl,
		AntplusDeviceTypeFitnessEquipment,
		AntplusDeviceTypeBloodPressure,
		AntplusDeviceTypeGeocacheNode,
		AntplusDeviceTypeLightElectricVehicle,
		AntplusDeviceTypeEnvSensor,
		AntplusDeviceTypeRacquet,
		AntplusDeviceTypeControlHub,
		AntplusDeviceTypeMuscleOxygen,
		AntplusDeviceTypeShifting,
		AntplusDeviceTypeBikeLightMain,
		AntplusDeviceTypeBikeLightShared,
		AntplusDeviceTypeExd,
		AntplusDeviceTypeBikeRadar,
		AntplusDeviceTypeBikeAero,
		AntplusDeviceTypeWeightScale,
		AntplusDeviceTypeHeartRate,
		AntplusDeviceTypeBikeSpeedCadence,
		AntplusDeviceTypeBikeCadence,
		AntplusDeviceTypeBikeSpeed,
		AntplusDeviceTypeStrideSpeedDistance,
	}
}
