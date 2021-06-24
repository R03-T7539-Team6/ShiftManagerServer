package main

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"
	"github.com/R03-T7539-Team6/ShiftManagerSerer/server"
)

func main() {
	db.Init()
	defer db.Close()
	server.Init()

}
