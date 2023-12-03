package handlers

import (
	"goClean/backend/models"
	"goClean/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStudent(s *gin.Context) {
	var stud *models.Student
	if err := s.ShouldBindJSON(&stud); err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Student{
		Name:     stud.Name,
		Password: stud.Password,
		RollNo:   stud.RollNo,
		Hostel:   stud.Hostel,
		RoomNo:   stud.RoomNo,
	}

	user.BeforeSave()
	user, err := user.SaveUser()
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, utils.Student)
}

func LoginStudent(s *gin.Context) {
	var stud *models.Student
	err := s.ShouldBindJSON(&stud)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(stud.Name, stud.Password)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.JSON(http.StatusOK, gin.H{"data": token})
}
