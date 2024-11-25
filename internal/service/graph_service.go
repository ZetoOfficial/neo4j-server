package service

import (
	"context"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
	"github.com/ZetoOfficial/neo4j-server/internal/repository"
)

type graphService struct {
	repo repository.Repository
}

func NewGraphService(repo repository.Repository) *graphService {
	return &graphService{repo: repo}
}

func (u *graphService) GetAllNodes(ctx context.Context) ([]models.GetAllNodesResponse, error) {
	return u.repo.GetAllNodes(ctx)
}

func (u *graphService) GetAllRelationships(ctx context.Context) ([]models.GetAllRelationshipsResponse, error) {
	return u.repo.GetAllRelationships(ctx)
}

func (u *graphService) GetNodeWithRelationships(ctx context.Context, nodeID int64) (models.GetNodeWithRelationshipsResponse, error) {
	return u.repo.GetNodeWithRelationships(ctx, nodeID)
}

func (u *graphService) AddNodeAndRelationships(ctx context.Context, req models.AddNodeAndRelationshipsRequest) error {
	return u.repo.AddNodeAndRelationships(ctx, req)
}

func (u *graphService) DeleteNodeAndRelationships(ctx context.Context, nodeID int64) error {
	return u.repo.DeleteNodeAndRelationships(ctx, nodeID)
}
