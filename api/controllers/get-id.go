package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetID(i Interface) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := strings.TrimSpace(c.Query("id"))
		id, err := strconv.Atoi(query)
		if err != nil {
			// handle error
			c.AbortWithStatusJSON(http.StatusBadRequest, "Please Provide Id and it should be Integer Type")

		} else {
			data, err1 := i.GetByID(id)
			if err1 != nil {
				c.JSON(http.StatusAccepted, *err1)
			} else {
				c.JSON(http.StatusOK, data)
			}
		}
	}
}
