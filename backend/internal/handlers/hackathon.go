package handlers

import (
	"net/http"
	"strconv"
	"teamfinder/backend/internal/models"
	"teamfinder/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type HackathonHandler struct {
	hackathonService *services.HackathonService
}

func NewHackathonHandler(hackathonService *services.HackathonService) *HackathonHandler {
	return &HackathonHandler{hackathonService: hackathonService}
}

func (h *HackathonHandler) GetAllHackathons(c *gin.Context) {
	hackathons, err := h.hackathonService.GetAllHackathons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hackathons"})
		return
	}
	c.JSON(http.StatusOK, hackathons)
}

func (h *HackathonHandler) GetHackathonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	hackathon, err := h.hackathonService.GetHackathonByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hackathon not found"})
		return
	}
	c.JSON(http.StatusOK, hackathon)
}

func (h *HackathonHandler) CreateHackathon(c *gin.Context) {
	var hackathon models.Hackathon
	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.hackathonService.CreateHackathon(&hackathon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hackathon"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Hackathon created successfully"})
}

func (h *HackathonHandler) UpdateHackathon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	var hackathon models.Hackathon
	if err := c.ShouldBindJSON(&hackathon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.hackathonService.UpdateHackathon(id, &hackathon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hackathon"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hackathon updated successfully"})
}

func (h *HackathonHandler) DeleteHackathon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	if err := h.hackathonService.DeleteHackathon(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hackathon"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Hackathon deleted successfully"})
}
