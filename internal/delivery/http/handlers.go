package http

import (
	"net/http"
	"strconv"

	"github.com/ZetoOfficial/neo4j-server/internal/models"
	"github.com/ZetoOfficial/neo4j-server/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.GraphService
}

func NewHandler(service service.GraphService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/nodes", h.GetAllNodes)
		api.GET("/nodes/:id", h.GetNodeWithRelations)

		authorized := api.Group("/")
		authorized.Use(AuthMiddleware())
		{
			authorized.POST("/nodes", h.CreateNodeAndRelationship)
			authorized.DELETE("/nodes/:id", h.DeleteNode)
		}
	}
}

func (h *Handler) GetAllNodes(c *gin.Context) {
	nodes, err := h.service.GetAllNodes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *Handler) GetNodeWithRelations(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	relations, err := h.service.GetNodeWithRelationships(c.Request.Context(), models.GetNodeWithRelationshipsRequest{NodeId: int64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relations)
}

func (h *Handler) CreateNodeAndRelationship(c *gin.Context) {
	var req models.CreateNodeAndRelationshipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	relations, err := h.service.CreateNodeAndRelationship(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, relations)
}

func (h *Handler) DeleteNode(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.service.DeleteNodeAndRelationships(c.Request.Context(), models.DeleteNodeAndRelationshipsRequest{NodeId: int64(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
