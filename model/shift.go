package model

import (
	"time"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Status value
/*
	None
	NotStarted
	Working
	FinalVersion
*/
type shiftStatus string

const (
	NoneShift    shiftStatus = "None"
	NotStarted   shiftStatus = "NotStarted"
	Working      shiftStatus = "Working"
	FinalVersion shiftStatus = "FinalVersion"
)

type Shift struct {
	WorkDate       time.Time `json:"work_date"`
	IsPaidHoliday  bool      `json:"is_paid_holiday"`
	AttendanceTime time.Time `json:"attendance_time"`
	LeavingTime    time.Time `json:"leaving_time"`
}

type ShiftRequest struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	LastUpdate time.Time `json:"last_upadate"`
	Shift      []Shift   `json:"shift_request"`
}

type ShiftSchedule struct {
	ID                   uint
	TargetDate           time.Time   `json:"target_date"`
	StartOfSchedule      time.Time   `json:"start_of_schedule"`
	EndOfSchedule        time.Time   `json:"end_of_schedule"`
	ShiftSchedulingState shiftStatus `json:"shift_state"`
	Shift                []Shift     `json:"shift_schedule"`
	Worker               []User      `json:"worker_list"`
}

// **************ShiftRequest
func (sr ShiftRequest) GetAll() ([]ShiftRequest, error) {
	db := db.GetDB()
	var srr []ShiftRequest

	if err := db.Find(&srr).Error; err != nil {
		return nil, err
	}
	return srr, nil
}

func (sr ShiftRequest) CreateShiftRequest(c *gin.Context) (ShiftRequest, error) {
	db := db.GetDB()
	if err := c.BindJSON(&sr); err != nil {
		return sr, err
	}

	if err := db.Create(&sr).Error; err != nil {
		return sr, err
	}
	return sr, nil
}

func (sr ShiftRequest) GetByUserId(user_id string) ([]ShiftRequest, error) {
	db := db.GetDB()
	var srr []ShiftRequest
	if err := db.Where("user_id = ?", user_id).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

func (sr ShiftRequest) DeleteByID(id string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Delete(&sr).Error; err != nil {
		return err
	}
	return nil
}

// **************ShiftSchedule
func (ss ShiftSchedule) CreateShiftSchedule(c *gin.Context) (ShiftSchedule, error) {
	db := db.GetDB()
	if err := c.BindJSON(&ss); err != nil {
		return ss, err
	}

	if err := db.Create(&ss).Error; err != nil {
		return ss, err
	}
	return ss, nil
}

func (ss ShiftSchedule) GetById(id string) (ShiftSchedule, error) {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Find(&ss).Error; err != nil {
		return ss, err
	}
	return ss, nil
}
