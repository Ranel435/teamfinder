package handlers

import (
	"log"
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
		log.Printf("ERROR: Failed to retrieve hackathons: %v", err)
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
		log.Printf("ERROR: Hackathon not found, id=%d: %v", id, err)
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
		log.Printf("ERROR: Failed to create hackathon '%s': %v", hackathon.Name, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hackathon"})
		return
	}

	log.Printf("INFO: Hackathon created successfully, id=%d, name='%s'", id, hackathon.Name)
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
		log.Printf("ERROR: Failed to update hackathon, id=%d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hackathon"})
		return
	}

	log.Printf("INFO: Hackathon updated successfully, id=%d, name='%s'", id, hackathon.Name)
	c.JSON(http.StatusOK, gin.H{"message": "Hackathon updated successfully"})
}

func (h *HackathonHandler) DeleteHackathon(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	if err := h.hackathonService.DeleteHackathon(id); err != nil {
		log.Printf("ERROR: Failed to delete hackathon, id=%d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hackathon"})
		return
	}

	log.Printf("WARN: Hackathon deleted, id=%d", id)
	c.JSON(http.StatusOK, gin.H{"message": "Hackathon deleted successfully"})
}
