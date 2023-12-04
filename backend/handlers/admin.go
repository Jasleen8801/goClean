package handlers

import (
	"fmt"
	"goClean/backend/models"
	"goClean/backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(a *gin.Context) {
	var adm *models.Admin
	if err := a.ShouldBindJSON(&adm); err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Admin{
		Name:     adm.Name,
		Password: adm.Password,
	}

	user.BeforeSave()
	_, err := user.SaveAdmin()
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.JSON(http.StatusOK, gin.H{"data": user, "message": "Admin registered successfully"})
}

func LoginAdmin(a *gin.Context) {
	var adm *models.Admin
	err := a.ShouldBindJSON(&adm)
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.AdminLoginCheck(adm.Name, adm.Password)
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.JSON(http.StatusOK, gin.H{"token": token, "message": "Admin logged in successfully"})
}

func AddHostel(a *gin.Context) {
	adminId, role, err := utils.ExtractTokenMetadata(a)
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(role)

	if role != utils.Admin {
		a.JSON(http.StatusBadRequest, gin.H{"error": "Not an admin"})
		return
	}

	admin, err := models.GetAdminById(adminId)
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var hostel *models.Hostel
	if err := a.ShouldBindJSON(&hostel); err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hostel.BeforeSave()
	hostel, err = admin.AddHostel(hostel)
	if err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.JSON(http.StatusOK, gin.H{"data": hostel, "message": "Hostel added successfully"})
}
