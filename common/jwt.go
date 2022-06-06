package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"ticket-backend/model"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.key"))

type UserClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type AdminClaims struct {
	AdminId uint `json:"admin_id"`
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
	claims := UserClaims{
		UserId: user.ID,
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

func ParseUserTokenString(tokenString string) (*jwt.Token, *UserClaims, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

func ReleaseAdminToken(admin model.Admin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 有效期7天
	claims := AdminClaims{
		AdminId: admin.ID,
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

func ParseAdminTokenString(tokenString string) (*jwt.Token, *AdminClaims, error) {
	claims := &AdminClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
