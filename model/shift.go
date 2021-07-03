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
	ID             uint      `json:"id" gorm:"unique"`
	UserID         string    `json:"user_id"`
	StoreID        string    `json:"store_id"`
	WorkDate       time.Time `json:"work_date" sql:"type:date"`
	IsPaidHoliday  bool      `json:"is_paid_holiday"`
	IsRequest      bool      `json:"is_request"`
	AttendanceTime time.Time `json:"attendance_time"`
	LeavingTime    time.Time `json:"leaving_time"`
	StartBreakTime time.Time `json:"start_break_time"`
	EndBreakTime   time.Time `json:"end_break_time"`
}

type ShiftRequest struct {
	gorm.Model
	UserID     string    `json:"user_id" gorm:"unique"`
	StoreID    string    `json:"store_id"`
	LastUpdate time.Time `json:"last_update"`
	Shift      []Shift   `json:"shifts" gorm:"foreignKey:WorkDate"`
}

type ShiftSchedule struct {
	gorm.Model
	StoreID              string      `json:"store_id"`
	TargetDate           time.Time   `json:"target_date" sql:"type:date"`
	StartOfSchedule      time.Time   `json:"start_of_schedule"`
	EndOfSchedule        time.Time   `json:"end_of_schedule"`
	ShiftSchedulingState shiftStatus `json:"shift_state"`
	Shift                []Shift     `json:"shifts" gorm:"foreignKey:UserID"`
	WorkerNum            uint        `json:"worker_num"`
}

// ******************Shift
//GetAllShift is get all shift
// func (sr Shift) GetAllShift() ([]Shift, error) {
// 	db := db.GetDB()
// 	var srr []Shift

// 	if err := db.Find(&srr).Error; err != nil {
// 		return nil, err
// 	}
// 	return srr, nil
// }

// CreateShift is create a shift
func (sr Shift) CreateShift(c *gin.Context) (Shift, error) {
	db := db.GetDB()
	if err := c.BindJSON(&sr); err != nil {
		return sr, err
	}

	if err := db.Create(&sr).Error; err != nil {
		return sr, err
	}
	return sr, nil
}

// GetByUserId is get all user shift
func (sr Shift) GetByUserId(user_id string) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("user_id = ?", user_id).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

// GetByUserIdAndRequest is get all user request shift
func (sr Shift) GetByUserIdAndIsRequest(user_id string, is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("user_id = ? AND is_request = ?", user_id, is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

// GetByWorkDate is get all work date shift.
func (sr Shift) GetByWorkDate(work_date time.Time) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("work_date = ?", work_date).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

// GetByWorkDateAndIsRequest is get all work date request shift.
func (sr Shift) GetByWorkDateAndIsRequest(work_date time.Time, is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("work_date = ? AND is_request = ?", work_date, is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

// GetByIsRequest is get all work date request shift.
func (sr Shift) GetByIsRequest(is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("is_request = ?", is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

// DeleteById is delete a shift by id
func (sr Shift) DeleteById(id string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Delete(&sr).Error; err != nil {
		return err
	}
	return nil
}

// ******************Shift

// **************ShiftRequest
// func (sr ShiftRequest) GetAll() ([]ShiftRequest, error) {
// 	db := db.GetDB()
// 	var srr []ShiftRequest

// 	if err := db.Find(&srr).Error; err != nil {
// 		return nil, err
// 	}
// 	return srr, nil
// }

// CreateShiftRequest is create a shift
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

// GetByUserId is get user request shift
// Simillar to Shift method (GetByUserIdAndIsRequest)
func (sr ShiftRequest) GetByUserId(user_id string) (ShiftRequest, error) {
	db := db.GetDB()

	var s []Shift
	var srr ShiftRequest
	if err := db.Where("user_id = ?", user_id).Find(&srr).Error; err != nil {
		return srr, err
	}

	if err := db.Where("user_id = ? AND is_request = true", user_id).Find(&s).Error; err != nil {
		return srr, err
	}

	// Join []User in ShiftRequest
	srr.Shift = s

	return srr, nil
}

// DeleteById is delete a shift request related user by id
func (sr ShiftRequest) DeleteById(id string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Delete(&sr).Error; err != nil {
		return err
	}
	return nil
}

// ****************ShfitRequest

// **************ShiftSchedule

// CreateShiftSchedule is a shift schedule
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

// GetById is get a shift schedule by id
func (ss ShiftSchedule) GetByStoreId(store_id string) (ShiftSchedule, error) {
	db := db.GetDB()
	var s []Shift
	if err := db.Where("store_id = ?", store_id).Find(&ss).Error; err != nil {
		return ss, err
	}

	if err := db.Where("store_id = ? AND is_request = false", store_id).Find(&s).Error; err != nil {
		return ss, err
	}
	return ss, nil
}

// ***************ShiftSchedule
