package model

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	StoreID       string          `json:"store_id" gorm:"unique"`
	Worker        []User          `json:"worker_lists" gorm:"foreignKey:UserID"`
	ShiftRequest  []ShiftRequest  `json:"shift_requests" gorm:"foreignKey:UserID"`
	ShiftSchedule []ShiftSchedule `json:"shift_schedules"`
}

/*************************************************
 *	specification;
 *	name 			= CreateStore
 *	Function 	= create store model in Store table
 *	note			= store table is related json
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http.request
 *  output    = Store: Store sturct
 * 						= error value
 *  end of specification;
**************************************************/
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

/*************************************************
 *	specification;
 *	name 			= GetByStoreID
 *	Function 	= Get a Store in store talbe with store id
 *	note			= store table is related json
 *						= get a user, request, schedule by store_id
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= c: *gin.Context http.request
 *  output    = Store: Store sturct
 * 						= error value
 *  end of specification;
**************************************************/
func (s Store) GetByStoreID(store_id string) (Store, error) {
	db := db.GetDB()
	var u []User
	var sr []ShiftRequest
	var ss []ShiftSchedule

	// Get Store struct
	if err := db.Where("store_id = ?", store_id).First(&s).Error; err != nil {
		return s, err
	}

	// Get users in store
	if err := db.Where("store_id = ?", store_id).Find(&u).Error; err != nil {
		return s, err
	}

	// Get shift requests in store
	if err := db.Where("store_id = ?", store_id).Find(&sr).Error; err != nil {
		return s, err
	}

	// Get shift schedule in store
	if err := db.Where("store_id = ?", store_id).Find(&ss).Error; err != nil {
		return s, err
	}

	// Give users shift request and schedule to store sturct
	s.Worker = u
	s.ShiftRequest = sr
	s.ShiftSchedule = ss

	return s, nil
}
