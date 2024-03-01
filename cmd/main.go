package main

import (
	"fmt"
	"invoice-service/pkg/adapters"
	"invoice-service/pkg/drivers"
	"log"
)

func main() {
	config, err := adapters.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	host := config.Host
	port := config.Port
	user := config.User
	password := config.Password
	dbname := config.Dbname

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	postgresAdapter, err := adapters.NewPostgresAdapter(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	r := drivers.SetupRouter(postgresAdapter)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
