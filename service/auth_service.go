package service

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	mySigningKey = []byte("s3cret")
)

type Token struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {

}

func GenerateJWT() (*Token, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usuario": "davi resio",
		"id": 5,
		"exp": time.Now().Add(time.Hour * 720).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("alguma coisa deu errado")
	}

	return &Token{Token: tokenString}, nil

}
