package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"ticket-backend/model"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.key"))

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//
//  ReleaseUserToken
//  @Description: 生成用户token
//  @param user
//  @return string
//  @return error
//
func ReleaseUserToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 有效期7天
	claims := Claims{
		Username: user.Name,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "northern air",
			Subject:   "user token",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}
