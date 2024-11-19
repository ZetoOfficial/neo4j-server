package repository

import (
	"context"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
)

type Repository interface {
	GetAllNodes(ctx context.Context) ([]models.GetAllNodesResponse, error)
	GetNodeWithRelationships(ctx context.Context, req models.GetNodeWithRelationshipsRequest) ([]models.GetNodeWithRelationshipsResponse, error)
	CreateNodeAndRelationship(ctx context.Context, req models.CreateNodeAndRelationshipRequest) (models.CreateNodeAndRelationshipResponse, error)
	DeleteNodeAndRelationships(ctx context.Context, req models.DeleteNodeAndRelationshipsRequest) error
}
