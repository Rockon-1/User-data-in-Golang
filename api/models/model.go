package models

type User struct {
	Id      int     `json:"ID"`
	Fname   string  `json:"First Name"`
	City    string  `json:"City"`
	Phone   uint64  `json:"Phone"`
	Height  float32 `json:"Height"`
	Married bool    `json:"Married status"`
}

type Error struct {
	Id      int
	Message string
}
