package main

import (
	"tansan/config"
	"tansan/database"
	"tansan/server"
)

func main() {
	cfg := config.GetConfig()
   
	db := database.NewPostgresDatabase(&cfg)
   
	server.NewEchoServer(&cfg, db.GetDb()).Start()
   }