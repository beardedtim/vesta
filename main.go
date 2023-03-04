package main

import (
	"fmt"
	"mckp/vesta/database"
	"mckp/vesta/env"
	mckttp "mckp/vesta/http"
)

func init() {
	database.Connect()
}

func main() {
	defer database.Disconnect()

	server := mckttp.New()

	server.Run(fmt.Sprintf(":%d", env.GetEnvInt("PORT")))
}
