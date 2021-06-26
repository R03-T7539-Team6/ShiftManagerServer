package entity

import "github.com/jinzhu/gorm"

type Store struct {
	gorm.Model
	StoreID       uint            `json:"store_id"`
	Worker        []User          `json:"worker_list"`
	ShiftRequest  []ShiftRequest  `json:"shift_request"`
	ShiftSchedule []ShiftSchedule `json:"shift_schedule"`
}
