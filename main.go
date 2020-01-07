package main

import (
	"os"

	"github.com/garaekz/goshort/url"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&url.URL{})

	return db
}
func main() {
	db := initDB()
	defer db.Close()

	urlAPI := initURLAPI(db)
	r := gin.Default()

	r.GET("/:code", urlAPI.FindByCode)
	v1 := r.Group("/api/v1/shorten")
	{
		v1.POST("/", urlAPI.Create)
	}

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
