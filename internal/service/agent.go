package service

import "awesomeProject2/internal/models"

type AgentRepo interface {
	Create(agent models.Agent) error
}
type Agent struct {
	repo AgentRepo
}

func (a *Agent) Create(agent models.Agent) error {
	return a.repo.Create(agent)
}
