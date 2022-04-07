package main

import (
	"github.com/rdenadai/gotodo/app"
	"github.com/rdenadai/gotodo/app/database"
)

func main() {
	database.StartDatabaseConnection()
	app.StartServer()
}
