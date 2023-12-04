package handlers

import (
	"goClean/backend/models"
	"goClean/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type feedbackInput struct {
	Feedback string `json:"feedback"`
}

func AddLog(c *gin.Context) {
	userId, role, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if role != utils.Student {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	stud, err := models.GetStudentById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var feedback feedbackInput
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Log := models.Log{
		StudentName: stud.Name,
		Hostel:      stud.Hostel,
		RoomNo:      stud.RoomNo,
		Floor:       stud.Floor,
		Feedback:    feedback.Feedback,
	}

	_, err = Log.SaveLog()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log added successfully"})
}
