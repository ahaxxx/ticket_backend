package main

import (
	"ticket-backed/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":3000")
}
