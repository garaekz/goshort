package main

import (
	"fmt"
	"log"
	"os"

	"github.com/garaekz/goshort/url"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func initDB() *gorm.DB {
	dbType := "postgres"

	if os.Getenv("APP_ENV") != "production" && os.Getenv("APP_ENV") != "travis" {
		dbType = "mysql"
	}

	db, err := gorm.Open(dbType, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Hubo un error al abrir DB")
		panic(err)
	}

	db.AutoMigrate(&url.URL{})

	return db
}
func main() {
	if os.Getenv("APP_ENV") != "production" && os.Getenv("APP_ENV") != "travis" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db := initDB()
	defer db.Close()

	urlAPI := initURLAPI(db)
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("./views/dist/**.html")

	r.Use(static.Serve("/", static.LocalFile("./views/dist", true)))
	r.GET("/:code", urlAPI.FindByCode)
	v1 := r.Group("/api/v1/shorten")
	{
		v1.POST("/", urlAPI.Create)
	}

	err := r.Run()
	if err != nil {
		fmt.Println("Hubo un error al iniciar el server")
		panic(err)
	}
}
