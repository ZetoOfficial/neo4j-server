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
		api.GET("/relationships", h.GetAllRelationships)
		api.GET("/nodes/:id", h.GetNodeWithRelationships)

		authorized := api.Group("/")
		authorized.Use(AuthMiddleware())
		{
			authorized.POST("/nodes", h.AddNodeAndRelationships)
			authorized.DELETE("/nodes/:id", h.DeleteNodeAndRelationships)
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

func (h *Handler) GetAllRelationships(c *gin.Context) {
	relationships, err := h.service.GetAllRelationships(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relationships)
}

func (h *Handler) GetNodeWithRelationships(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	relations, err := h.service.GetNodeWithRelationships(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relations)
}

func (h *Handler) AddNodeAndRelationships(c *gin.Context) {
	var req models.AddNodeAndRelationshipsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	err := h.service.AddNodeAndRelationships(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) DeleteNodeAndRelationships(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.service.DeleteNodeAndRelationships(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
