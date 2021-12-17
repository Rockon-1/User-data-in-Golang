package main

import (
	"PROJECT1/api/controllers"
	"PROJECT1/dataservices"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var i dataservices.Interface = &dataservices.Client{}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/all", controllers.Getall(i))
	r.GET("/id", controllers.GetID(i))
	r.GET("/ListOfID", controllers.GetByListOfID(i))

	r.Run()
}
