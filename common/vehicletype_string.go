// Code generated by "stringer -type=VehicleType"; DO NOT EDIT

package common

import "fmt"

const _VehicleType_name = "BusTramTrolleySubway"

var _VehicleType_index = [...]uint8{0, 3, 7, 14, 20}

func (i VehicleType) String() string {
	if i < 0 || i >= VehicleType(len(_VehicleType_index)-1) {
		return fmt.Sprintf("VehicleType(%d)", i)
	}
	return _VehicleType_name[_VehicleType_index[i]:_VehicleType_index[i+1]]
}