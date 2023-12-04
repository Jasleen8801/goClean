package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func TokenValid(c *gin.Context) error {
	token := c.Request.Header.Get("Authorization")
	// fmt.Println(token)
	if token == "" {
		return jwt.ErrSignatureInvalid
	}

	if len(strings.Split(token, " ")) != 2 {
		return jwt.ErrSignatureInvalid
	}

	token = strings.Split(token, " ")[1]

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return err
}

func ExtractTokenMetadata(c *gin.Context) (uint, UserRole, error) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		return 0, "", jwt.ErrSignatureInvalid
	}

	if len(strings.Split(token, " ")) != 2 {
		return 0, "", jwt.ErrSignatureInvalid
	}

	token = strings.Split(token, " ")[1]

	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if !ok || !tokenClaims.Valid {
		return 0, "", jwt.ErrSignatureInvalid
	}

	id, err := strconv.Atoi(fmt.Sprintf("%.0f", claims["user_id"]))
	if err != nil {
		return 0, "", err
	}

	role := UserRole(claims["role"].(string))

	return uint(id), role, nil
}
