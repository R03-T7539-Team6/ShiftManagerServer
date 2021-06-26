package entity

import "errors"

// Group value
/*
	None,										// デフォルト値
	SystemAdmin							// システム管理者（データ編集用)
	SuperUser							　// 権限ユーザ（店舗責任者）
	NormalUser							// 一般ユーザ
	ForTimeRecordTerminal		// 勤怠登録/店舗端末用
*/

type Group string

const (
	None                  Group = "None"
	SystemAdmin           Group = "SystemAdmin"
	SuperUser             Group = "SuperUser"
	NormalUser            Group = "NormalUser"
	ForTimeRecordTerminal Group = "ForTimeRecordTerminal"
)

// 使わないかも知れない
func (s *Group) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	*s = Group(string(asBytes))
	return nil
}

type UserGroup struct {
	ID    uint  `json:"id"`
	Group Group `json:"group"`
}
