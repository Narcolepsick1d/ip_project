package models

type ClientCart struct {
	Id       int `json:"id"`
	ClientId int `json:"client_id"`
	HouseId  int `json:"house_id"`
}
