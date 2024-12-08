package util

import (
	"github.com/imniynaiy/ticket-system/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(config.GlobalConfig.Server.AuthSalt+password))
}
