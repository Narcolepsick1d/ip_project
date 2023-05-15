package models

type Owner struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Address string `json:"address"`
}
