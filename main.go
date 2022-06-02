package main

import (
	"os"
	"ticket-backend/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
