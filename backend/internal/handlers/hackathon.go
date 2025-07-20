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

// GetAllHackathons godoc
// @Summary      Get all hackathons
// @Description  Retrieve a list of all available hackathons
// @Tags         Hackathons
// @Produce      json
// @Success      200 {array} models.Hackathon "List of hackathons"
// @Failure      500 {object} map[string]string "Failed to retrieve hackathons"
// @Router       /api/hackathons [get]
func (h *HackathonHandler) GetAllHackathons(c *gin.Context) {
	hackathons, err := h.hackathonService.GetAllHackathons()
	if err != nil {
		log.Printf("ERROR: Failed to retrieve hackathons: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hackathons"})
		return
	}
	c.JSON(http.StatusOK, hackathons)
}

// GetHackathonByID godoc
// @Summary      Get hackathon by ID
// @Description  Retrieve a specific hackathon by its ID
// @Tags         Hackathons
// @Produce      json
// @Param        id path int true "Hackathon ID"
// @Success      200 {object} models.Hackathon "Hackathon details"
// @Failure      400 {object} map[string]string "Invalid hackathon ID"
// @Failure      404 {object} map[string]string "Hackathon not found"
// @Router       /api/hackathons/{id} [get]
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

// CreateHackathon godoc
// @Summary      Create a new hackathon
// @Description  Create a new hackathon with the provided details
// @Tags         Hackathons
// @Accept       json
// @Produce      json
// @Param        hackathon body models.Hackathon true "Hackathon data (ID will be ignored)"
// @Success      201 {object} map[string]interface{} "Hackathon created successfully"
// @Failure      400 {object} map[string]string "Invalid request data"
// @Failure      500 {object} map[string]string "Failed to create hackathon"
// @Router       /api/hackathons [post]
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

// UpdateHackathon godoc
// @Summary      Update a hackathon
// @Description  Update an existing hackathon by ID
// @Tags         Hackathons
// @Accept       json
// @Produce      json
// @Param        id path int true "Hackathon ID"
// @Param        hackathon body models.Hackathon true "Updated hackathon data"
// @Success      200 {object} map[string]string "Hackathon updated successfully"
// @Failure      400 {object} map[string]string "Invalid hackathon ID or request data"
// @Failure      500 {object} map[string]string "Failed to update hackathon"
// @Router       /api/hackathons/{id} [put]
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

// DeleteHackathon godoc
// @Summary      Delete a hackathon
// @Description  Delete an existing hackathon by ID
// @Tags         Hackathons
// @Produce      json
// @Param        id path int true "Hackathon ID"
// @Success      200 {object} map[string]string "Hackathon deleted successfully"
// @Failure      400 {object} map[string]string "Invalid hackathon ID"
// @Failure      500 {object} map[string]string "Failed to delete hackathon"
// @Router       /api/hackathons/{id} [delete]
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
