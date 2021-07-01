package model

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	StoreID       string          `json:"store_id"`
	Worker        []User          `json:"worker_list" gorm:"foreignKey:UserID"`
	ShiftRequest  []ShiftRequest  `json:"shift_request" gorm:"foreignKey:UserID"`
	ShiftSchedule []ShiftSchedule `json:"shift_schedule"`
}

func (s Store) CreateStore(c *gin.Context) (Store, error) {
	db := db.GetDB()

	if err := c.BindJSON(&s); err != nil {
		return s, err
	}
	if err := db.Create(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}

func (s Store) GetByID(id string) (Store, error) {
	db := db.GetDB()
	if err := db.Where("store_id = ?", id).First(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}
