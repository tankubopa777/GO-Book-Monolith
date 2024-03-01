package main

import (
	"tansan/config"
	"tansan/database"
	"tansan/user/entities"
)

func main() {
 cfg := config.GetConfig()

 db := database.NewPostgresDatabase(&cfg)

 userMigrate(db)
}

func userMigrate(db database.Database) {
 db.GetDb().Migrator().CreateTable(&entities.User{})
 db.GetDb().CreateInBatches([]entities.User{
  {Amount: 1},
  {Amount: 2},
  {Amount: 2},
  {Amount: 5},
  {Amount: 3},
 }, 10)
}