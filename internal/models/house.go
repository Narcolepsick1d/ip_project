package models

import "time"

type House struct {
	Id        int       `json:"id,omitempty"`
	Address   string    `json:"address"`
	OwnerId   int       `json:"owner_id"`
	AgentId   int       `json:"agent_id"`
	BuildDate time.Time `json:"build_date"`
	Price     float64   `json:"price"`
}
