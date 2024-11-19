package models

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Request/Response models for GetAllNodes
type GetAllNodesResponse struct {
	NodeId     int64    `json:"nodeId"`
	NodeLabels []string `json:"nodeLabels"`
}

// Request/Response models for GetNodeWithRelationships
type GetNodeWithRelationshipsRequest struct {
	NodeId int64 `json:"nodeId"`
}

type GetNodeWithRelationshipsResponse struct {
	Node         neo4j.Node         `json:"node"`
	Relationship neo4j.Relationship `json:"relationship"`
	TargetNode   neo4j.Node         `json:"targetNode"`
}

// Модель для свойств узла
type NodeProperties struct {
	City       string `json:"city" validate:"required,max=100"` // Обязательно, максимум 100 символов
	ScreenName string `json:"screen_name" validate:"required"`  // Обязательно
	Sex        int    `json:"sex" validate:"gte=0,lte=2"`       // Значение от 0 до 2
	Name       string `json:"name" validate:"required,max=100"` // Обязательно, максимум 100 символов
	ID         int    `json:"id" validate:"required,min=1"`     // Обязательно, минимум 1
}

// Модель для свойств связи
type RelationshipProperties struct {
	Since                int    `json:"since" validate:"gte=1900,lte=2100"` // Год в диапазоне 1900-2100
	RelationshipStrength string `json:"relationshipStrength" validate:"max=50"`
}

// Request/Response models for CreateNodeAndRelationship
type CreateNodeAndRelationshipRequest struct {
	NodeProps         NodeProperties         `json:"nodeProps"`
	RelationshipType  string                 `json:"relationshipType" validate:"required,oneof=SUBSCRIBES FOLLOWS FRIENDS"`
	RelationshipProps RelationshipProperties `json:"relationshipProps"`
	RelatedNodeProps  NodeProperties         `json:"relatedNodeProps"`
}

type CreateNodeAndRelationshipResponse struct {
	CreatedNode         neo4j.Node         `json:"createdNode"`
	CreatedRelationship neo4j.Relationship `json:"createdRelationship"`
	CreatedRelatedNode  neo4j.Node         `json:"createdRelatedNode"`
}

// Request models for Delete operations
type DeleteNodeAndRelationshipsRequest struct {
	NodeId int64 `json:"nodeId"`
}

type DeleteGraphSegmentRequest struct {
	NodeId int64 `json:"nodeId"`
}
