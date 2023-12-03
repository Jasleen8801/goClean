package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserRole string

const (
	Student UserRole = "student"
	Hostel  UserRole = "hostel"
	Admin   UserRole = "admin"
)

func GenerateToken(user_id uint, role UserRole) (string, error) {
	// Generates a JWT token
	tokenLifespan, error := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))
	if error != nil {
		return "", error
	}

	claims := jwt.MapClaims{} // jwt.MapClaims is a map[string]interface{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	// Extracts the user_id from the JWT token
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		return 0, jwt.ErrSignatureInvalid
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	user_id, err := strconv.Atoi(fmt.Sprintf("%.0f", claims["user_id"]))
	if err != nil {
		return 0, err
	}

	return uint(user_id), nil
}
