package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(config string) *mysqlDatabase {
	mysql := &mysqlDatabase{
		db:            &gorm.DB{},
		configuration: config,
	}
	return mysql
}

type mysqlDatabase struct {
	db            *gorm.DB
	configuration string
}

func (m *mysqlDatabase) connect() error {
	var connectToDatabaseError error
	m.db, connectToDatabaseError = gorm.Open(mysql.Open(m.configuration), &gorm.Config{})

	if connectToDatabaseError != nil {
		log.Fatal("Error connecting to database")
	}

	return nil
}
