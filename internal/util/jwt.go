package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/imniynaiy/ticket-system/internal/config"
	"github.com/imniynaiy/ticket-system/internal/log"
)

var mySigningKey = []byte(config.GlobalConfig.Server.JwtSigningKey)

type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func CreateJWT(username string) (token string, err error) {
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "flome",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(mySigningKey)
}

func VerifyJWT(token string) error {
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, keyFunc)
	if err != nil {
		return err
	} else if claims, ok := t.Claims.(*CustomClaims); ok {
		log.Info("user authenticated: ", log.String("username", claims.Username))
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return errors.New("unknown claims type")
	}
	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return mySigningKey, nil
}
