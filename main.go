package main

import (
	"AdcashTask/Database"
	"AdcashTask/Models"
	"AdcashTask/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	//database.DB, err = gorm.Open("mysql", database.DbURL(database.BuildDBConfig()))
	database.DB, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection Successfully Opened ")
	database.DB.AutoMigrate(&Models.Category{}, &Models.Product{})

	fmt.Println("Database Migrated")
}

func main() {

	initDatabase()
	defer database.DB.Close()
	r := Routes.SetupRouter()
	//running
	_ = r.Run()

	defer database.DB.Close()

}
