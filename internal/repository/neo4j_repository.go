package repository

import (
	"context"
	"fmt"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jStorage struct {
	driver neo4j.DriverWithContext
}

func NewNeo4jStorage(uri, username, password string) (*Neo4jStorage, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	if err = driver.VerifyConnectivity(ctx); err != nil {
		driver.Close(ctx)
		return nil, err
	}
	return &Neo4jStorage{driver: driver}, nil
}

func (s *Neo4jStorage) GetAllNodes(ctx context.Context) ([]models.GetAllNodesResponse, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (n) RETURN id(n) AS nodeId, labels(n) AS nodeLabels", nil)
	if err != nil {
		return nil, err
	}

	var responses []models.GetAllNodesResponse
	for result.Next(ctx) {
		record := result.Record()
		nodeId, ok := record.Get("nodeId")
		if !ok {
			continue
		}
		nodeLabelsInterface, ok := record.Get("nodeLabels")
		if !ok {
			continue
		}
		nodeLabelsSlice, ok := nodeLabelsInterface.([]interface{})
		if !ok {
			continue
		}
		nodeLabels := make([]string, len(nodeLabelsSlice))
		for i, v := range nodeLabelsSlice {
			if label, ok := v.(string); ok {
				nodeLabels[i] = label
			} else {
				return nil, fmt.Errorf("expected string label but got %T", v)
			}
		}

		responses = append(responses, models.GetAllNodesResponse{
			NodeId:     nodeId.(int64),
			NodeLabels: nodeLabels,
		})
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("get all nodes: %w", err)
	}

	return responses, nil
}

func (s *Neo4jStorage) GetNodeWithRelationships(ctx context.Context, req models.GetNodeWithRelationshipsRequest) ([]models.GetNodeWithRelationshipsResponse, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, "MATCH (n)-[r]->(m) WHERE id(n) = $nodeId RETURN n, r, m", map[string]interface{}{
		"nodeId": req.NodeId,
	})
	if err != nil {
		return nil, err
	}

	var responses []models.GetNodeWithRelationshipsResponse
	for result.Next(ctx) {
		record := result.Record()
		n, _ := record.Get("n")
		r, _ := record.Get("r")
		m, _ := record.Get("m")
		responses = append(responses, models.GetNodeWithRelationshipsResponse{
			Node:         n.(neo4j.Node),
			Relationship: r.(neo4j.Relationship),
			TargetNode:   m.(neo4j.Node),
		})
	}

	if err = result.Err(); err != nil {
		return nil, err
	}

	return responses, nil
}

// CreateNodeAndRelationship - создание узла и связи
// CreateNodeAndRelationship - создание узла и связи
func (s *Neo4jStorage) CreateNodeAndRelationship(ctx context.Context, req models.CreateNodeAndRelationshipRequest) (models.CreateNodeAndRelationshipResponse, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Формируем запрос Cypher с использованием RelationshipType
	cypherQuery := `
		CREATE (n:Label $nodeProps)-[r:%s $relProps]->(m:AnotherLabel $relatedNodeProps)
		RETURN n, r, m
	`
	cypherQuery = fmt.Sprintf(cypherQuery, req.RelationshipType)

	// Преобразуем свойства для передачи в запрос
	nodeProps := map[string]interface{}{
		"city":        req.NodeProps.City,
		"screen_name": req.NodeProps.ScreenName,
		"sex":         req.NodeProps.Sex,
		"name":        req.NodeProps.Name,
		"id":          req.NodeProps.ID,
	}

	relatedNodeProps := map[string]interface{}{
		"city":        req.RelatedNodeProps.City,
		"screen_name": req.RelatedNodeProps.ScreenName,
		"sex":         req.RelatedNodeProps.Sex,
		"name":        req.RelatedNodeProps.Name,
		"id":          req.RelatedNodeProps.ID,
	}

	relProps := map[string]interface{}{
		"since":                req.RelationshipProps.Since,
		"relationshipStrength": req.RelationshipProps.RelationshipStrength,
	}

	// Выполняем запрос
	result, err := session.Run(ctx, cypherQuery, map[string]interface{}{
		"nodeProps":        nodeProps,
		"relProps":         relProps,
		"relatedNodeProps": relatedNodeProps,
	})
	if err != nil {
		return models.CreateNodeAndRelationshipResponse{}, fmt.Errorf("create node and relationship: %w", err)
	}

	// Извлекаем результаты
	if result.Next(ctx) {
		record := result.Record()
		n, _ := record.Get("n")
		r, _ := record.Get("r")
		m, _ := record.Get("m")
		return models.CreateNodeAndRelationshipResponse{
			CreatedNode:         n.(neo4j.Node),
			CreatedRelationship: r.(neo4j.Relationship),
			CreatedRelatedNode:  m.(neo4j.Node),
		}, nil
	}

	return models.CreateNodeAndRelationshipResponse{}, fmt.Errorf("no record found")
}

// DeleteNodeAndRelationships - удаление узла и его связей
func (s *Neo4jStorage) DeleteNodeAndRelationships(ctx context.Context, req models.DeleteNodeAndRelationshipsRequest) error {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, "MATCH (n) WHERE id(n) = $nodeId DETACH DELETE n", map[string]interface{}{
		"nodeId": req.NodeId,
	})
	return err
}
