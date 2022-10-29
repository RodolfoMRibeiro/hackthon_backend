package main

import (
	"hackthon/adapter/env"
	"hackthon/pluggy"
)

func main() {
	pluggy.GenerateAccessToken()
	env.Load()
	// db.StartDatabase()
	// seed.SeedDatabase()
	// servidor := server.CreateServer()
	// router.Avaible(servidor.GetServerEngine())
	// servidor.Run()
}
