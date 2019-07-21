package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "akshat"
	dbname   = "world_events"
)

var DB gorm.DB

var gormOpen = gorm.Open

func InitDB() error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err := gormOpen("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer DB.Close()

	// TODO: Auto migrate all the things
	// DB.AutoMigrate()

	fmt.Println("Successfully connected!")

	return nil
}
