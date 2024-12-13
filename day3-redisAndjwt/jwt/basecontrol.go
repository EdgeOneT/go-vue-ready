package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var secretKey = []byte("mysecretKey")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成Token
func generateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "myApp",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析Token
func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func main() {
	token, err := generateToken("john")
	if err != nil {
		log.Fatal("Error generating token:", err)
	}
	fmt.Println("Generated Token:", token)

	claims, err := parseToken(token)
	if err != nil {
		log.Fatal("Error parsing token:", err)
	}

	fmt.Println("Parsed claims:", claims.Username)
}
