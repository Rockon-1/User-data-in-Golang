package controllers

import (
	"PROJECT1/dataservices"
)

type Interface interface {
	Getall() ([]dataservices.User, *dataservices.Error)
	GetByID(i int) (*dataservices.User, *dataservices.Error)
}
