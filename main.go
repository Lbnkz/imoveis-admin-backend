package main

import (
	"imoveis-back/database"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.StartDB_SQL()
	db := database.GetDatabase_SQL()
	print(db)
}
