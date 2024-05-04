package main

import (
	"example.com/crud-api-with-postgres/db"
	"example.com/crud-api-with-postgres/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
