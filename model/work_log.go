package model

import (
	"time"

	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type WorkLog struct {
	gorm.Model
	UserID         string      `json:"user_id"`
	AttendanceTime time.Time   `json:"attendance_time"`
	LeavingTime    time.Time   `json:"leaving_time"`
	BreakTime      []time.Time `json:"break_time"`
}

func (w WorkLog) Create(c *gin.Context) (WorkLog, error) {
	db := db.GetDB()
	if err := c.BindJSON(&w); err != nil {
		return w, nil
	}

	if err := db.Create(&w).Error; err != nil {
		return w, err
	}
	return w, nil
}

func (w WorkLog) GetByUserID(user_id string) ([]WorkLog, error) {
	db := db.GetDB()
	var ww []WorkLog
	if err := db.Where("user_id = ?", user_id).Find(&ww).Error; err != nil {
		return nil, err
	}
	return ww, nil
}
