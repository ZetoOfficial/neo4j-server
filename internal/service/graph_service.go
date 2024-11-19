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

func (u *graphService) GetNodeWithRelationships(ctx context.Context, req models.GetNodeWithRelationshipsRequest) ([]models.GetNodeWithRelationshipsResponse, error) {
	return u.repo.GetNodeWithRelationships(ctx, req)
}

func (u *graphService) CreateNodeAndRelationship(ctx context.Context, req models.CreateNodeAndRelationshipRequest) (models.CreateNodeAndRelationshipResponse, error) {
	return u.repo.CreateNodeAndRelationship(ctx, req)
}

func (u *graphService) DeleteNodeAndRelationships(ctx context.Context, req models.DeleteNodeAndRelationshipsRequest) error {
	return u.repo.DeleteNodeAndRelationships(ctx, req)
}
