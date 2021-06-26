package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type WorkLog struct {
	gorm.Model
	UserID         uint      `json:"user_id"`
	AttendanceTime time.Time `json:"attendance_time"`
	LeavingTime    time.Time `json:"leaving_time"`
	BreakTime      time.Time `json:"break_time"`
}
