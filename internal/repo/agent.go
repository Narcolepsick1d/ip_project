package repo

import (
	"awesomeProject2/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type AgentPostgres struct {
	db *pgxpool.Pool
}

func NewAgentPostgres(db *pgxpool.Pool) *AgentPostgres {
	return &AgentPostgres{db}
}
func (r *AgentPostgres) Create(agent models.Agent) error {

	_, err := r.db.Exec(context.Background(), "INSERT INTO agents (user_id,phone) VALUES ($1,$2)", agent.Id, agent.Phone)
	log.Println(err)
	return err
}
