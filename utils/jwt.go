package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims((jwt.SigningMethodHS256), jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims")
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userIdClaim, ok := claims["userId"]
	if !ok {
		fmt.Println("userId claim missing in token")
		return 0, errors.New("userId claim missing in token")
	}

	userIdFloat, ok := userIdClaim.(float64)
	if !ok {
		fmt.Println("userId claim is not a float64")
		return 0, errors.New("userId claim is not a float64")
	}

	userId := int64(userIdFloat)
	fmt.Println(userId)

	return userId, nil

}
