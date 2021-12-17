package dataservices

type Client struct{}
type Interface interface {
	Getall() ([]User, *Error)
	GetByID(i int) (*User, *Error)
}
