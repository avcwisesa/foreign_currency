package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	c "github.com/avcwisesa/foreign_currency/controller"
	d "github.com/avcwisesa/foreign_currency/database"
	m "github.com/avcwisesa/foreign_currency/model"
)

func main() {

	client, _ := gorm.Open("sqlite3", "db.sqlite")
	database := d.New(client)

	database.Migrate(&m.ExchangeRate{})
	database.Migrate(&m.TrackedExchange{})

	controller := c.New(database)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/exchangeRate", controller.GetExchangeRateHist)
	r.GET("/trackedExchange", controller.GetTrackedExchanges)
	r.POST("/exchangeRate/add", controller.AddExchangeRate)
	r.POST("/trackedExchange/add", controller.AddTrackedExchange)
	r.DELETE("/trackedExchange/delete", controller.DeleteTrackedExchange)

	r.Run() // listen and serve on 0.0.0.0:8080
}