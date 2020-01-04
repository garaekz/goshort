package main

import (
	"os"

	"github.com/jinzhu/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&url.Url{})

	return db
}
func main() {
	db := initDB()
	defer db.Close()

	urlAPI := InitURLApi(db)
}
