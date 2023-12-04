package handlers

import (
	"goClean/backend/models"
	"goClean/backend/utils"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHostel(h *gin.Context) {
	var host *models.Hostel
	err := h.ShouldBindJSON(&host)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host.Name = html.EscapeString(host.Name)

	token, err := models.HostelLoginCheck(host.Name, host.Password)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{"token": token, "message": "Hostel logged in successfully"})
}

func ClearAllLogs(h *gin.Context) {
	hostId, role, err := utils.ExtractTokenMetadata(h)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if role != utils.Hostel {
		h.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	host, err := models.GetHostelById(hostId)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = host.ClearAllLogs()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{"message": "All logs cleared successfully"})
}
