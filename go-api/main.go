package main

import (
	"goapi/db"

	"goapi/server"
)

func main() {

	// db
	db.InitDB()
	defer db.CloseDB()

	// start server
	server.Run()
}
