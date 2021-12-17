package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getall(i Interface) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := i.Getall()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Some Issue Happened !!! Sorry for Inconvenigence")
		} else {
			c.JSON(http.StatusOK, data)
		}
	}
}
