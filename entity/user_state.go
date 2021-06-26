package entity

import "errors"

// Status value
/*
	Normal          	// 通常
	InLeaveOfAbsence	// 休職中
	Retired         	// 退職済み
	NotHired        	// 未採用
	Others          	// その他
*/

type Status string

const (
	Normal           Status = "Normal"
	InLeaveOfAbsence Status = "InLeaveOfAbsence"
	Retired          Status = "Retired"
	NotHired         Status = "NotHired"
	Others           Status = "Others"
)

// 使わないかも知れない
func (s *Status) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	*s = Status(string(asBytes))
	return nil
}

type UserState struct {
	ID     uint   `json:"id"`
	Status Status `json:"status"`
}
