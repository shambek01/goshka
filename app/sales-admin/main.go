package main

import (
	"fmt"
	"goshka/app/sales-admin/commands"
	"goshka/business/data/schema"
	"goshka/foundation/database"
	"log"
)

func main() {
	//commands.GenKey()
	commands.GenToken()
	cfg := database.Config{
		User:       "postgres",
		Password:   "123",
		Host:       "0.0.0.0",
		Name:       "ardan_labs",
		DisableTLS: true,
	}

	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalln("connecting database")
	}
	defer db.Close()

	if err := schema.Migrate(db); err != nil {
		log.Fatalln("migrating database")
	}

	fmt.Println("migrations complete")

	if err := schema.Seed(db); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("seed data complete")
}
