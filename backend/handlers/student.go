package handlers

import (
	"fmt"
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
		Floor:    stud.Floor,
	}

	user.BeforeSave()
	_, err := user.SaveUser()
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.JSON(http.StatusOK, gin.H{"data": user, "message": "Student registered successfully"})
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

	s.JSON(http.StatusOK, gin.H{"token": token, "message": "Student logged in successfully"})
}

func GetStudentById(s *gin.Context) {
	userId, _, err := utils.ExtractTokenMetadata(s)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := models.GetStudentById(userId)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.JSON(http.StatusOK, gin.H{"data": student, "message": "Student fetched successfully"})
}

func GetStudentsbyRoomID(s *gin.Context) {
	roomId := s.Param("room_no")
	if roomId == "" {
		s.JSON(http.StatusBadRequest, gin.H{"error": "Room number required"})
		return
	}

	fmt.Println(roomId)

	students, err := models.GetStudentsbyRoomID(roomId)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.JSON(http.StatusOK, gin.H{"data": students, "message": "Students fetched successfully"})
}

func UpdateRoomStatus(s *gin.Context) {
	roomId := s.Param("room_no")
	if roomId == "" {
		s.JSON(http.StatusBadRequest, gin.H{"error": "Room number required"})
		return
	}

	err := models.UpdateRoomCleaned(roomId)
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.JSON(http.StatusOK, gin.H{"message": "Room status updated successfully"})
}
