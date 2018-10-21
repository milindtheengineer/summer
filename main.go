package main

import (
	"fmt"
	"log"
	"os"
	"summer/brain"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")
	err := godotenv.Load("secrets.env")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	brain.ServeRequests()
}
