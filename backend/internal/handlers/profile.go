package handlers

import (
	"log"
	"net/http"
	"strconv"
	"teamfinder/backend/internal/models"
	"teamfinder/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfilesByHackathonID(c *gin.Context) {
	hackathonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	profiles, err := h.profileService.GetProfilesByHackathonID(hackathonID)
	if err != nil {
		log.Printf("ERROR: Failed to retrieve profiles for hackathon %d: %v", hackathonID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profiles"})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	hackathonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile.HackathonID = hackathonID
	if profile.Status == "" {
		profile.Status = "active"
	}

	id, err := h.profileService.CreateProfile(&profile)
	if err != nil {
		log.Printf("ERROR: Failed to create profile for user %d, hackathon %d: %v", profile.UserID, hackathonID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	log.Printf("INFO: Profile created successfully, id=%d, user_id=%d, hackathon_id=%d, role='%s'",
		id, profile.UserID, hackathonID, profile.DesiredRole)
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Profile created successfully"})
}

func (h *ProfileHandler) GetProfileByID(c *gin.Context) {
	hackathonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	profileID, err := strconv.Atoi(c.Param("profile_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	profile, err := h.profileService.GetProfileByID(hackathonID, profileID)
	if err != nil {
		log.Printf("ERROR: Profile not found, hackathon_id=%d, profile_id=%d: %v", hackathonID, profileID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	hackathonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	profileID, err := strconv.Atoi(c.Param("profile_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.profileService.UpdateProfile(hackathonID, profileID, &profile); err != nil {
		log.Printf("ERROR: Failed to update profile, hackathon_id=%d, profile_id=%d: %v", hackathonID, profileID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	log.Printf("INFO: Profile updated successfully, hackathon_id=%d, profile_id=%d, role='%s'",
		hackathonID, profileID, profile.DesiredRole)
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	hackathonID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hackathon ID"})
		return
	}

	profileID, err := strconv.Atoi(c.Param("profile_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile ID"})
		return
	}

	if err := h.profileService.DeleteProfile(hackathonID, profileID); err != nil {
		log.Printf("ERROR: Failed to delete profile, hackathon_id=%d, profile_id=%d: %v", hackathonID, profileID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete profile"})
		return
	}

	log.Printf("WARN: Profile deleted, hackathon_id=%d, profile_id=%d", hackathonID, profileID)
	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}
