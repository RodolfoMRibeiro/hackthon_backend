package db

import (
	"fmt"
	"hackthon/adapter/env"
	"log"

	"gorm.io/gorm"
)

var db *gorm.DB

func GetGormDB() *gorm.DB {
	return db
}

func StartDatabase() {
	databaseConfiguration := createDatabaseStringConfig()
	database := NewMysql(databaseConfiguration)

	if err := database.connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	loadMigrations(database.db)

	db = database.db

	fmt.Println("Connected to Database sucessfully")
}

func createDatabaseStringConfig() string {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",

		env.Mysql.USER,
		env.Mysql.PASSWORD,
		env.Mysql.HOST,
		env.Mysql.PORT,
		env.Mysql.DB,
		env.Mysql.ADDITIONAL_CONFIGS,
	)

	return databaseStringConfig
}

func loadMigrations(db *gorm.DB) {

}
