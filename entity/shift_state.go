package entity

import "errors"

// Status value
/*
	None
	NotStarted
	Working
	FinalVersion
*/

type ShiftStatus string

const (
	NoneShift    ShiftStatus = "None"
	NotStarted   ShiftStatus = "NotStarted"
	Working      ShiftStatus = "Working"
	FinalVersion ShiftStatus = "FinalVersion"
)

func (s *ShiftStatus) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	*s = ShiftStatus(string(asBytes))
	return nil
}

type ShiftState struct {
	ID     uint        `json:"id"`
	Status ShiftStatus `json:"status"`
}
