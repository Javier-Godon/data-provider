package main

import (
	"log"

	"data-provider/framework"
)

func main() {
	framework.ReadConfig()

	framework.InitDatabase()

	defer framework.DB.Close()

	log.Println("Application started successfully")
}
