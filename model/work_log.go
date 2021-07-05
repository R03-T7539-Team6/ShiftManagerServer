package model

import (
	"time"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type WorkLog struct {
	gorm.Model
	UserID         string    `json:"user_id"`
	AttendanceTime time.Time `json:"attendance_time"`
	LeavingTime    time.Time `json:"leaving_time"`
	StartBreakTime time.Time `json:"start_break_time"`
	EndBreakTime   time.Time `json:"end_break_time"`
}

/*************************************************
 *	specification;
 *	name 			= CreateLog
 *	Function 	= Create row WorkLog table
 *	note			= worklog table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http request
 *  output    = WorkLog: Worklog struct
 * 						= error value
 *  end of specification;
**************************************************/
func (w WorkLog) CreateLog(c *gin.Context) (WorkLog, error) {
	db := db.GetDB()
	if err := c.BindJSON(&w); err != nil {
		return w, nil
	}

	if err := db.Create(&w).Error; err != nil {
		return w, err
	}
	return w, nil
}

/*************************************************
 *	specification;
 *	name 			= UpdateByID
 *	Function 	= Update selected row by id
 *	note			= worklog table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http request
 * 						= id: Worklog id
 *  output    = WorkLog: Worklog struct
 * 						= error value
 *  end of specification;
**************************************************/
func (w WorkLog) UpdateByID(id string, c *gin.Context) (WorkLog, error) {
	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&w).Error; err != nil {
		return w, err
	}
	if err := c.BindJSON(&w); err != nil {
		return w, err
	}
	db.Save(&w)
	return w, nil
}

/*************************************************
 *	specification;
 *	name 			= GetByUserID
 *	Function 	= Get all User log
 *	note			= worklog table is related json
							= id is user_id,  search table by user id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= user_id: string
 *  output    = []WorkLog: Worklog struct array
 * 						= error value
 *  end of specification;
**************************************************/
func (w WorkLog) GetByUserID(user_id string) ([]WorkLog, error) {
	db := db.GetDB()
	var ww []WorkLog
	if err := db.Where("user_id = ?", user_id).Find(&ww).Error; err != nil {
		return nil, err
	}
	return ww, nil
}
