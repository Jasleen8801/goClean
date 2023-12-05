package handlers

import (
	"fmt"
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

func authenticateHostel(h *gin.Context) (*models.Hostel, error) {
	hostId, role, err := utils.ExtractTokenMetadata(h)
	if err != nil {
		return &models.Hostel{}, err
	}

	if role != utils.Hostel {
		return &models.Hostel{}, fmt.Errorf("Unauthorized")
	}

	host, err := models.GetHostelById(hostId)
	if err != nil {
		return &models.Hostel{}, err
	}

	return host, nil
}

func ClearAllLogs(h *gin.Context) {
	host, err := authenticateHostel(h)
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

func GetAllUncleanedRooms(h *gin.Context) {
	host, err := authenticateHostel(h)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rooms, err := host.GetAllUncleanedRoomNos()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func GetAllCleanedRooms(h *gin.Context) {
	host, err := authenticateHostel(h)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rooms, err := host.GetAllCleanedRoomNos()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func GetAllHostelLogs(h *gin.Context) {
	host, err := authenticateHostel(h)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logs, err := host.GetAllLogs()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{"logs": logs})
}
