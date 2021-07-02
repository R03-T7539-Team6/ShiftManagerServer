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

// CreateLog is create a log
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

// UpdateByID is update log by id
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

// GetByUserID is get all user log
func (w WorkLog) GetByUserID(user_id string) ([]WorkLog, error) {
	db := db.GetDB()
	var ww []WorkLog
	if err := db.Where("user_id = ?", user_id).Find(&ww).Error; err != nil {
		return nil, err
	}
	return ww, nil
}
