package main

import (
	"hackthon/adapter/env"
	"hackthon/pluggy"
)

func do() {

}

func main() {
	env.Load()
	pluggy.GenerateAccessToken()
	pluggy.CreateItem()
	// db.StartDatabase()
	// seed.SeedDatabase()
	// servidor := server.CreateServer()
	// router.Avaible(servidor.GetServerEngine())
	// servidor.Run()
}
