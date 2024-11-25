package models

import "github.com/neo4j/neo4j-go-driver/v5/neo4j"

type GetAllNodesResponse struct {
	ID    int64   `json:"id"`
	Label string  `json:"label"`
	Name  *string `json:"name,omitempty"`
}

type GetAllRelationshipsResponse struct {
	StartNodeID      int64      `json:"start_node_id"`
	RelationshipType string     `json:"relationship_type"`
	EndNodeID        int64      `json:"end_node_id"`
	EndNode          neo4j.Node `json:"end_node"`
}
type GetNodeWithRelationshipsResponse struct {
	Node          Node           `json:"node"`
	Relationships []Relationship `json:"relationships"`
}

type AddNodeAndRelationshipsRequest struct {
	Node          Node           `json:"node"`
	Relationships []Relationship `json:"relationships"`
}
