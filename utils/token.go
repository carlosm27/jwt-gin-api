package utils

import (
	"errors"
	"fmt"
	"github.com/carlosm27/jwtGinApi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"strings"
	"time"
	
)

var privateKey = []byte(os.Getenv("API_SECRET"))

func GenerateToken(user models.User) (string, error) {

	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(privateKey)

}


func ValidateToken (c *gin.Context) error {
	token, err := GetToken(c)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token provided")
}

func CurrentUser(c *gin.Context) (models.User, error) {
	err := ValidateToken(c)
    if err != nil {
        return models.User{}, err
    }
    token, _ := GetToken(c)
    claims, _ := token.Claims.(jwt.MapClaims)
    userId := uint(claims["id"].(float64))

    user, err := models.GetUserById(userId)
    if err != nil {
        return models.User{}, err
    }
    return user, nil
}


func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return privateKey, nil
    })
    return token, err
}


func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
