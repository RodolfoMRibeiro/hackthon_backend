package main

import (
	"hackthon/adapter/db"
	"hackthon/adapter/db/seed"
	"hackthon/adapter/env"
	"hackthon/adapter/router"
	"hackthon/adapter/server"
)

func main() {
	env.Load()
	db.StartDatabase()
	seed.SeedDatabase()
	servidor := server.CreateServer()
	router.Avaible(servidor.GetServerEngine())
	servidor.Run()
}
