package repository

import (
	"service/internal/config"
	"service/internal/infrastructure/database"
	"service/internal/model"
)

type AgentReportRepository struct {
	Config *config.Configuration
	Db     *database.DB
}

func NewAgentReportRepository(
	cfg *config.Configuration,
	db *database.DB,
) *AgentReportRepository {
	return &AgentReportRepository{
		Config: cfg,
		Db:     db,
	}
}

func (s *AgentReportRepository) Create(agentReport *model.AgentReport) (bool, error) {
	result := s.Db.Connection.Create(agentReport) // pass pointer of data to Create
	if result != nil {
		return true, nil
	}
	return false, nil
}
