package main

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/R03-T7539-Team6/ShiftManagerSerer/model"
	"github.com/R03-T7539-Team6/ShiftManagerSerer/server"
)

func main() {
	db.Init(false,
		&model.User{},
		&model.WorkLog{},
		&model.Shift{},
		&model.Store{},
	)
	defer db.Close()
	server.Init()
}
