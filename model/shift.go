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

/*************************************************
 *	specification;
 *	name 			= CreateShift
 *	Function 	= create a shift in shift table
 *	note			= shift table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http.request
 *  output    = Shift: Shift sturct
 * 						= error value
 *  end of specification;
**************************************************/
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

/*************************************************
 *	specification;
 *	name 			= GetByUserId
 *	Function 	= Get shifts in shift table
 *	note			= shift table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= user_id: string
 *  output    = []Shift: Shift sturct array
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) GetByUserId(user_id string) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("user_id = ?", user_id).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByUserIdAndIsRequest
 *	Function 	= Get shifts in shift table
 *	note			= shift table is related json
 *            = can get a request shift
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= user_id: string
 * 						= is_request: bool
 *  output    = []Shift: Shift sturct array
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) GetByUserIdAndIsRequest(user_id string, is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("user_id = ? AND is_request = ?", user_id, is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByUserIDandWorkDate
 *	Function 	= Get shifts in shift table
 *	note			= shift table is related json
 *            = can get a specific date shift
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= work_date: time.Time UTC value
 *  output    = []Shift: Shift sturct array
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) GetByUserIDandWorkDate(user_id string, work_date string) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("user_id = ? AND work_date = ?", user_id, work_date).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByWorkDateAndIsRequest
 *	Function 	= get shifts in shift table
 *	note			= shift table is related json
 * 						= can get specific workdate, and request shifts
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= work_date: time.Time UTC value
 * 						= is_request: bool
 *  output    = []Shift: Shift sturct array
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) GetByWorkDateAndIsRequest(work_date string, is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("work_date = ? AND is_request = ?", work_date, is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByIsRequest
 *	Function 	= get shifts in shift table
 *	note			= shift table is related json
 * 						= can get request shifts
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= is_request: bool
 *  output    = []Shift: Shift sturct array
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) GetByIsRequest(is_request bool) ([]Shift, error) {
	db := db.GetDB()
	var srr []Shift
	if err := db.Where("is_request = ?", is_request).Find(&srr).Error; err != nil {
		return srr, err
	}
	return srr, nil
}

/*************************************************
 *	specification;
 *	name 			= UpdateShift
 *	Function 	= Update a shift in shift table
 *	note			= shift table is related json
 * 						= id is not user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string
 * 						= c: *gin.Context
 *  output    = Shift: Shift sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (sr Shift) UpdateShift(id string, c *gin.Context) (Shift, error) {
	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&sr).Error; err != nil {
		return sr, err
	}
	if err := c.BindJSON(&sr); err != nil {
		return sr, err
	}
	return sr, nil
}

/*************************************************
 *	specification;
 *	name 			= DeleteById
 *	Function 	= Delete a shift in shift table
 *	note			= shift table is related json
 * 						= id is not user_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= id: string
 *  output    = error value
 *  end of specification;
**************************************************/
func (sr Shift) DeleteById(id string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Delete(&sr).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************
 *	specification;
 *	name 			= CreateShiftRequest
 *	Function 	= Create a shift request model in shiftRequest table
 *	note			= shiftRequest table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context
 *  output    = ShiftRequest: ShiftRequest value
 * 						= error: error value
 *  end of specification;
**************************************************/
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

/*************************************************
 *	specification;
 *	name 			= GetByUserId
 *	Function 	= get a shift request model in shiftRequest table by user_id
 *	note			= shiftRequest table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= user_id: string
 *  output    = ShiftRequest: ShiftRequest value
 * 						= error: error value
 *  end of specification;
**************************************************/
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

/*************************************************
 *	specification;
 *	name 			= DeleteById
 *	Function 	= delete a shift request model in shiftRequest table by shift id
 *	note			= shiftRequest table is related json
 * 						= id is not user id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= user_id: string
 *  output    = id: string
 *  end of specification;
**************************************************/
func (sr ShiftRequest) DeleteById(id string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", id).Delete(&sr).Error; err != nil {
		return err
	}
	return nil
}

/*************************************************
 *	specification;
 *	name 			= CreateShiftSchedule
 *	Function 	= Create a shift schedule model in shiftSchedule table
 *	note			= shiftSchedule table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http.Request
 *  output    = ShiftSchedule: ShiftSchedule struct
 * 						= error: error value
 *  end of specification;
**************************************************/
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

/*************************************************
 *	specification;
 *	name 			= GetByStoreId
 *	Function 	= Get a shift schedule model in shiftSchedule table by store Id
 *	note			= shiftSchedule table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= store_id: string
 *  output    = ShiftSchedule: ShiftSchedule struct
 * 						= error: error value
 *  end of specification;
**************************************************/
func (ss ShiftSchedule) GetByStoreId(store_id string) (ShiftSchedule, error) {
	db := db.GetDB()
	var s []Shift
	if err := db.Where("store_id = ?", store_id).Find(&ss).Error; err != nil {
		return ss, err
	}

	if err := db.Where("store_id = ? AND is_request = false", store_id).Find(&s).Error; err != nil {
		return ss, err
	}

	ss.Shift = s
	return ss, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByStoreIdAndTargetDate
 *	Function 	= Get a shift schedule model in shiftSchedule table by store Id
 *	note			= shiftSchedule table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= store_id: string
 *  output    = ShiftSchedule: ShiftSchedule struct
 * 						= error: error value
 *  end of specification;
**************************************************/
func (ss ShiftSchedule) GetByStoreIdAndTargetDate(store_id string, target_date string) (ShiftSchedule, error) {
	db := db.GetDB()
	var s []Shift
	if err := db.Where("store_id = ? AND target_date = ?", store_id, target_date).Find(&ss).Error; err != nil {
		return ss, err
	}

	if err := db.Where("store_id = ? AND target_date = ? AND is_request = false", store_id, target_date).Find(&s).Error; err != nil {
		return ss, err
	}

	ss.Shift = s
	return ss, nil
}
