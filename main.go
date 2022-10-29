package main

import (
	"fmt"
	"hackthon/adapter/env"
	"hackthon/pluggy"
)

func do() {

}

//item
//8b520971-16d2-4b7c-8385-29fa94e30938
//account
//cd57ad15-ed0e-4374-81f2-417ab55895c4
//transaction
//fb5d644c-1498-4194-baa3-103b9a271a5e
func main() {
	env.Load()
	pluggy.GenerateApiKey()
	fmt.Println(pluggy.ListAccounts("693372d5-83be-4eb9-8a09-ea95d6304963"))
	// db.StartDatabase()
	// seed.SeedDatabase()
	// servidor := server.CreateServer()
	// router.Avaible(servidor.GetServerEngine())
	// servidor.Run()
}
