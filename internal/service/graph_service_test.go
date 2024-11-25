package service

import (
	"context"
	"errors"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
	"github.com/ZetoOfficial/neo4j-server/internal/repository/mocks"
)

func TestGraphService_GetAllNodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	expectedNodes := []models.GetAllNodesResponse{
		{
			ID:    1,
			Label: "Person",
			Name:  stringPointer("Alice"),
		},
		{
			ID:    2,
			Label: "Company",
			Name:  stringPointer("Acme Corp"),
		},
	}

	mockRepo.EXPECT().GetAllNodes(ctx).Return(expectedNodes, nil)

	nodes, err := service.GetAllNodes(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedNodes, nodes)
}

func TestGraphService_GetAllRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	expectedRelationships := []models.GetAllRelationshipsResponse{
		{
			StartNodeID:      1,
			RelationshipType: "FRIEND",
			EndNodeID:        2,
			EndNode:          mockNeo4jNode(2, "Person", "Bob"),
		},
	}

	mockRepo.EXPECT().GetAllRelationships(ctx).Return(expectedRelationships, nil)

	relationships, err := service.GetAllRelationships(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedRelationships, relationships)
}

func TestGraphService_GetNodeWithRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	nodeID := int64(1)
	expectedResult := models.GetNodeWithRelationshipsResponse{
		Node: models.Node{
			ID:         1,
			Label:      "Person",
			Name:       stringPointer("Alice"),
			ScreenName: stringPointer("alice123"),
			Sex:        intPointer(1),
			City:       stringPointer("New York"),
		},
		Relationships: []models.Relationship{
			{
				Type:      "FRIEND",
				EndNodeID: 2,
			},
		},
	}

	mockRepo.EXPECT().GetNodeWithRelationships(ctx, nodeID).Return(expectedResult, nil)

	result, err := service.GetNodeWithRelationships(ctx, nodeID)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestGraphService_AddNodeAndRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	req := models.AddNodeAndRelationshipsRequest{
		Node: models.Node{
			ID:    1,
			Label: "Person",
			Name:  stringPointer("Alice"),
		},
		Relationships: []models.Relationship{
			{
				Type:      "FRIEND",
				EndNodeID: 2,
			},
		},
	}

	mockRepo.EXPECT().AddNodeAndRelationships(ctx, req).Return(nil)

	err := service.AddNodeAndRelationships(ctx, req)
	assert.NoError(t, err)
}

func TestGraphService_DeleteNodeAndRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	nodeID := int64(1)

	mockRepo.EXPECT().DeleteNodeAndRelationships(ctx, nodeID).Return(nil)
	err := service.DeleteNodeAndRelationships(ctx, nodeID)
	assert.NoError(t, err)
}

func TestGraphService_GetAllNodes_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewGraphService(mockRepo)

	ctx := context.Background()
	expectedError := errors.New("failed to get all nodes")

	mockRepo.EXPECT().GetAllNodes(ctx).Return(nil, expectedError)

	nodes, err := service.GetAllNodes(ctx)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, nodes)
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

func mockNeo4jNode(id int64, label, name string) neo4j.Node {
	props := map[string]interface{}{
		"name": name,
	}
	return neo4j.Node{
		Id:     id,
		Labels: []string{label},
		Props:  props,
	}
}
