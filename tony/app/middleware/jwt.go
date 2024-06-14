package middleware

import (
	"errors"
	"log/slog"
	"time"

	"github.com/HsiaoCz/monster-clone/leaf/conf"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// define our claims
type myClaims struct {
	UserID  primitive.ObjectID `json:"userID"`
	Email   string             `json:"email"`
	IsAdmin bool               `json:"isAdmin"`
	jwt.StandardClaims
}

// define expire time
const TokenExpireDuration = time.Hour * 24 * 3

var mySecret = []byte(conf.Conf.App.JWTSecret)

// GenToken generate jwt token
func GenToken(userID primitive.ObjectID, email string, isAdmin bool) (string, error) {
	claims := myClaims{
		userID,
		email,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "hotel-hsiaol1",
		},
	}
	// use jwt generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		slog.Error("token failed", "err", err)
		return "", err
	}
	return token, nil
}

// ParseToken parse JWT
func ParseToken(tokenString string) (*myClaims, error) {
	// parse
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
