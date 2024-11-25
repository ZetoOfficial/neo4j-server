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

	query := `
		MATCH (n)
		RETURN n.id AS id, labels(n) AS label, n.name AS name, n.screen_name AS screen_name, n.sex AS sex, n.city AS city
	`
	result, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	var responses []models.GetAllNodesResponse
	for result.Next(ctx) {
		record := result.Record()
		id, _ := record.Get("id")
		label, _ := record.Get("label")
		name, _ := record.Get("name")

		var namePtr *string

		if name != nil {
			nameStr := name.(string)
			namePtr = &nameStr
		}

		responses = append(responses, models.GetAllNodesResponse{
			ID:    id.(int64),
			Label: label.([]interface{})[0].(string),
			Name:  namePtr,
		})
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("get all nodes: %w", err)
	}

	return responses, nil
}

func (s *Neo4jStorage) GetAllRelationships(ctx context.Context) ([]models.GetAllRelationshipsResponse, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (n)-[r]->(m)
		RETURN n.id AS start_node_id, type(r) AS relationship_type, m.id AS end_node_id, m
	`
	result, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, err
	}

	var responses []models.GetAllRelationshipsResponse
	for result.Next(ctx) {
		record := result.Record()
		startNodeID, _ := record.Get("start_node_id")
		relationshipType, _ := record.Get("relationship_type")
		endNodeID, _ := record.Get("end_node_id")
		endNode, _ := record.Get("m")

		responses = append(responses, models.GetAllRelationshipsResponse{
			StartNodeID:      startNodeID.(int64),
			RelationshipType: relationshipType.(string),
			EndNodeID:        endNodeID.(int64),
			EndNode:          endNode.(neo4j.Node),
		})
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("get all relationships: %w", err)
	}

	return responses, nil
}

func (s *Neo4jStorage) GetNodeWithRelationships(ctx context.Context, nodeID int64) (models.GetNodeWithRelationshipsResponse, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
        MATCH (n {id: $id})-[r]-(m)
        RETURN n, type(r) AS relationship_type, m
    `
	result, err := session.Run(ctx, query, map[string]interface{}{"id": nodeID})
	if err != nil {
		return models.GetNodeWithRelationshipsResponse{}, err
	}

	var node models.Node
	relationshipsMap := make(map[int64]models.Relationship)

	for result.Next(ctx) {
		record := result.Record()
		n, _ := record.Get("n")
		relType, _ := record.Get("relationship_type")
		m, _ := record.Get("m")

		// Process the main node
		if node.ID == 0 {
			nodeProps := n.(neo4j.Node).Props
			node = models.Node{
				ID:    nodeProps["id"].(int64),
				Label: n.(neo4j.Node).Labels[0],
			}
			if name, ok := nodeProps["name"].(string); ok {
				node.Name = &name
			}
			if screenName, ok := nodeProps["screen_name"].(string); ok {
				node.ScreenName = &screenName
			}
			if sex, ok := nodeProps["sex"].(int64); ok {
				sexInt := int(sex)
				node.Sex = &sexInt
			}
			if city, ok := nodeProps["city"].(string); ok {
				node.City = &city
			}
		}

		// Process the related node and relationship
		endNode := m.(neo4j.Node)
		endNodeID := endNode.Props["id"].(int64)
		if _, exists := relationshipsMap[endNodeID]; !exists {
			relationship := models.Relationship{
				Type:      relType.(string),
				EndNodeID: endNodeID,
			}
			relationshipsMap[endNodeID] = relationship
		}
	}

	if err = result.Err(); err != nil {
		return models.GetNodeWithRelationshipsResponse{}, fmt.Errorf("get node with relationships: %w", err)
	}

	// Convert map to slice
	relationships := make([]models.Relationship, 0, len(relationshipsMap))
	for _, rel := range relationshipsMap {
		relationships = append(relationships, rel)
	}

	return models.GetNodeWithRelationshipsResponse{Node: node, Relationships: relationships}, nil
}

func (s *Neo4jStorage) AddNodeAndRelationships(ctx context.Context, req models.AddNodeAndRelationshipsRequest) error {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MERGE (n:User {id: $id})
		SET n.name = $name, n.screen_name = $screen_name, n.sex = $sex, n.city = $city
		UNWIND $relationships AS rel
		MATCH (m:User {id: rel.end_node_id})
		MERGE (n)-[:FOLLOW]->(m)
	`
	params := map[string]interface{}{
		"id":            req.Node.ID,
		"name":          req.Node.Name,
		"screen_name":   req.Node.ScreenName,
		"sex":           req.Node.Sex,
		"city":          req.Node.City,
		"relationships": req.Relationships,
	}

	_, err := session.Run(ctx, query, params)
	if err != nil {
		return fmt.Errorf("add node and relationships: %w", err)
	}

	return nil
}

func (s *Neo4jStorage) DeleteNodeAndRelationships(ctx context.Context, nodeID int64) error {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MATCH (n {id: $id})
		DETACH DELETE n
	`
	_, err := session.Run(ctx, query, map[string]interface{}{"id": nodeID})
	if err != nil {
		return fmt.Errorf("delete node and relationships: %w", err)
	}

	return nil
}
