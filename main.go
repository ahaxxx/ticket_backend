package main

import (
	"ticket-backed/database"
	"ticket-backed/router"
)

func main() {
	db := database.InitDB()
	defer db.Close()
	r := router.NewRouter()
	r.Run(":3000")
}
