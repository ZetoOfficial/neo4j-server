package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
	"github.com/ZetoOfficial/neo4j-server/internal/service"
	"github.com/ZetoOfficial/neo4j-server/internal/service/mocks"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestGetAllNodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	expectedNodes := []models.GetAllNodesResponse{
		{NodeId: 1, NodeLabels: []string{"Node1"}},
		{NodeId: 2, NodeLabels: []string{"Node2"}},
	}

	mockRepo.EXPECT().GetAllNodes(ctx).Return(expectedNodes, nil)
	nodes, err := service.GetAllNodes(ctx)
	assert.NoError(t, err)
	assert.Equal(t, expectedNodes, nodes)
}

func TestGetAllNodes_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	mockError := errors.New("repository error")

	mockRepo.EXPECT().GetAllNodes(ctx).Return(nil, mockError)
	nodes, err := service.GetAllNodes(ctx)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)
	assert.Nil(t, nodes)
}

func TestGetNodeWithRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	request := models.GetNodeWithRelationshipsRequest{
		NodeId: 1,
	}

	expectedResponse := []models.GetNodeWithRelationshipsResponse{
		{
			Node:         createMockNode(1, []string{"Node1"}),
			Relationship: createMockRelationship("CONNECTED_TO"),
			TargetNode:   createMockNode(2, []string{"Node2"}),
		},
	}

	mockRepo.EXPECT().GetNodeWithRelationships(ctx, request).Return(expectedResponse, nil)
	response, err := service.GetNodeWithRelationships(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestCreateNodeAndRelationship(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	request := models.CreateNodeAndRelationshipRequest{
		NodeProps: models.NodeProperties{
			City:       "New York",
			ScreenName: "user123",
			Sex:        1,
			Name:       "John Doe",
			ID:         3,
		},
		RelationshipType: "CONNECTED_TO",
		RelationshipProps: models.RelationshipProperties{
			Since:                2020,
			RelationshipStrength: "Strong",
		},
		RelatedNodeProps: models.NodeProperties{
			City:       "Los Angeles",
			ScreenName: "user456",
			Sex:        2,
			Name:       "Jane Smith",
			ID:         4,
		},
	}

	expectedResponse := models.CreateNodeAndRelationshipResponse{
		CreatedNode:         createMockNode(3, []string{"Node3"}),
		CreatedRelationship: createMockRelationship("CONNECTED_TO"),
		CreatedRelatedNode:  createMockNode(4, []string{"Node4"}),
	}

	mockRepo.EXPECT().CreateNodeAndRelationship(ctx, request).Return(expectedResponse, nil)
	response, err := service.CreateNodeAndRelationship(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func TestDeleteNodeAndRelationships(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	request := models.DeleteNodeAndRelationshipsRequest{
		NodeId: 1,
	}

	mockRepo.EXPECT().DeleteNodeAndRelationships(ctx, request).Return(nil)
	err := service.DeleteNodeAndRelationships(ctx, request)
	assert.NoError(t, err)
}

func TestDeleteNodeAndRelationships_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockGraphService(ctrl)
	service := service.NewGraphService(mockRepo)
	ctx := context.Background()

	request := models.DeleteNodeAndRelationshipsRequest{
		NodeId: 1,
	}

	mockError := errors.New("delete error")

	mockRepo.EXPECT().DeleteNodeAndRelationships(ctx, request).Return(mockError)
	err := service.DeleteNodeAndRelationships(ctx, request)
	assert.Error(t, err)
	assert.Equal(t, mockError, err)
}

func createMockNode(id int64, labels []string) neo4j.Node {
	return neo4j.Node{
		Id:     id,
		Labels: labels,
	}
}

func createMockRelationship(relType string) neo4j.Relationship {
	return neo4j.Relationship{
		Type: relType,
	}
}
