package repository

import (
	"context"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
)

type Repository interface {
	GetAllNodes(ctx context.Context) ([]models.GetAllNodesResponse, error)
	GetAllRelationships(ctx context.Context) ([]models.GetAllRelationshipsResponse, error)
	GetNodeWithRelationships(ctx context.Context, nodeID int64) (models.NodeWithRelationships, error)
	AddNodeAndRelationships(ctx context.Context, req models.AddNodeAndRelationshipsRequest) error
	DeleteNodeAndRelationships(ctx context.Context, nodeID int64) error
}
