// Code generated by "enumer -type=RunType -yaml -json -text -transform=upper -output=runtype_strings.go"; DO NOT EDIT.

package runtype

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _RunTypeName = "NONEPHYSICSTECHNICALPEDESTALPULSERCALIBRATIONCOSMIC"

var _RunTypeIndex = [...]uint8{0, 4, 11, 20, 28, 34, 45, 51}

const _RunTypeLowerName = "nonephysicstechnicalpedestalpulsercalibrationcosmic"

func (i RunType) String() string {
	if i < 0 || i >= RunType(len(_RunTypeIndex)-1) {
		return fmt.Sprintf("RunType(%d)", i)
	}
	return _RunTypeName[_RunTypeIndex[i]:_RunTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _RunTypeNoOp() {
	var x [1]struct{}
	_ = x[NONE-(0)]
	_ = x[PHYSICS-(1)]
	_ = x[TECHNICAL-(2)]
	_ = x[PEDESTAL-(3)]
	_ = x[PULSER-(4)]
	_ = x[CALIBRATION-(5)]
	_ = x[COSMIC-(6)]
}

var _RunTypeValues = []RunType{NONE, PHYSICS, TECHNICAL, PEDESTAL, PULSER, CALIBRATION, COSMIC}

var _RunTypeNameToValueMap = map[string]RunType{
	_RunTypeName[0:4]:        NONE,
	_RunTypeLowerName[0:4]:   NONE,
	_RunTypeName[4:11]:       PHYSICS,
	_RunTypeLowerName[4:11]:  PHYSICS,
	_RunTypeName[11:20]:      TECHNICAL,
	_RunTypeLowerName[11:20]: TECHNICAL,
	_RunTypeName[20:28]:      PEDESTAL,
	_RunTypeLowerName[20:28]: PEDESTAL,
	_RunTypeName[28:34]:      PULSER,
	_RunTypeLowerName[28:34]: PULSER,
	_RunTypeName[34:45]:      CALIBRATION,
	_RunTypeLowerName[34:45]: CALIBRATION,
	_RunTypeName[45:51]:      COSMIC,
	_RunTypeLowerName[45:51]: COSMIC,
}

var _RunTypeNames = []string{
	_RunTypeName[0:4],
	_RunTypeName[4:11],
	_RunTypeName[11:20],
	_RunTypeName[20:28],
	_RunTypeName[28:34],
	_RunTypeName[34:45],
	_RunTypeName[45:51],
}

// RunTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RunTypeString(s string) (RunType, error) {
	if val, ok := _RunTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _RunTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to RunType values", s)
}

// RunTypeValues returns all values of the enum
func RunTypeValues() []RunType {
	return _RunTypeValues
}

// RunTypeStrings returns a slice of all String values of the enum
func RunTypeStrings() []string {
	strs := make([]string, len(_RunTypeNames))
	copy(strs, _RunTypeNames)
	return strs
}

// IsARunType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i RunType) IsARunType() bool {
	for _, v := range _RunTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for RunType
func (i RunType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for RunType
func (i *RunType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RunType should be a string, got %s", data)
	}

	var err error
	*i, err = RunTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for RunType
func (i RunType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for RunType
func (i *RunType) UnmarshalText(text []byte) error {
	var err error
	*i, err = RunTypeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for RunType
func (i RunType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for RunType
func (i *RunType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = RunTypeString(s)
	return err
}