package main

import (
	"log"
	"os"

	"github.com/mickey-mickser/mini-bank/internal/migrations"
)

func main() {
	dbURL := "postgres://bank:bank@localhost:5437/bank?sslmode=disable"
	if len(os.Args) < 2 {
		log.Fatal("use: migrate up | down")
	}

	switch os.Args[1] {
	case "up":
		if err := migrations.Up(dbURL); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := migrations.Down(dbURL); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command")
	}

	log.Println("done")
	log.Println("migrations applied")
}
