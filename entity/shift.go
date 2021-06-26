package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Shift struct {
	gorm.Model
	WorkDate       time.Time `json:"work_date"`
	IsPaidHoliday  bool      `json:"is_paid_holiday"`
	AttendanceTime time.Time `json:"attendance_time"`
	LeavingTime    time.Time `json:"leaving_time"`
}

type ShiftRequest struct {
	UserID     uint      `json:"user_id"`
	LastUpdate time.Time `json:"last_upadate"`
	Shift      []Shift   `json:"shift_request"`
}

type ShiftSchedule struct {
	TargetDate           time.Time  `json:"target_date"`
	StartOfSchedule      time.Time  `json:"start_of_schedule"`
	EndOfSchedule        time.Time  `json:"end_of_schedule"`
	ShiftSchedulingState ShiftState `json:"shift_state"`
	Shift                []Shift    `json:"shift_schedule"`
	Worker               []User     `json:"worker_list"`
}
