package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func GenerateJwtToken(userId uint, email string, role string) (string, error) {
	lifespan, err := strconv.Atoi(viper.GetString("TOKEN_LIFESPAN_HOUR"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("API_SECRET")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, string, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("API_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, "", err
		}
		role := fmt.Sprintf("%s", claims["role"])
		return uint(uid), role, nil
	}
	return 0, "", nil
}
