// generated by jsonenums -type=VehicleType; DO NOT EDIT

package common

import (
	"encoding/json"
	"fmt"
)

var (
	_VehicleTypeNameToValue = map[string]VehicleType{
		"Bus":     Bus,
		"Tram":    Tram,
		"Trolley": Trolley,
		"Subway":  Subway,
	}

	_VehicleTypeValueToName = map[VehicleType]string{
		Bus:     "Bus",
		Tram:    "Tram",
		Trolley: "Trolley",
		Subway:  "Subway",
	}
)

func init() {
	var v VehicleType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_VehicleTypeNameToValue = map[string]VehicleType{
			interface{}(Bus).(fmt.Stringer).String():     Bus,
			interface{}(Tram).(fmt.Stringer).String():    Tram,
			interface{}(Trolley).(fmt.Stringer).String(): Trolley,
			interface{}(Subway).(fmt.Stringer).String():  Subway,
		}
	}
}

// MarshalJSON is generated so VehicleType satisfies json.Marshaler.
func (r VehicleType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _VehicleTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid VehicleType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so VehicleType satisfies json.Unmarshaler.
func (r *VehicleType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("VehicleType should be a string, got %s", data)
	}
	v, ok := _VehicleTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid VehicleType %q", s)
	}
	*r = v
	return nil
}
