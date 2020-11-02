package config

import (
	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

// var AuthConfig2 = middleware.JWTConfig{
// 	Claims:     &jwtCustomClaims{},
// 	SigningKey: signingKey,
// }

// var AuthConfig = middleware.KeyAuthConfig{
// 	KeyLookup: "header:hoge",
// 	Validator: func(key string, c echo.Context) (bool, error) {
// 		return key == "hoge", nil
// 	},
// }
