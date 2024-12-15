package util

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/imniynaiy/ticket-system/internal/config"
	"github.com/imniynaiy/ticket-system/internal/database"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

const expirationTime = 24 * time.Hour
const cost = 10

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(config.GlobalConfig.Server.AuthSalt+password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func CompareHashAndPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(config.GlobalConfig.Server.AuthSalt+password))
}

func getKey(token string) string {
	return "session:" + token
}

func GenTokenAndStoreInRedis(us *model.UserSession) (string, error) {

	token := uuid.New().String()
	key := getKey(token)
	sessionJSON, err := json.Marshal(us)
	if err != nil {
		return "", err
	}
	err = database.GlobalRedis.SetEx(context.Background(), key, sessionJSON, expirationTime).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyTokenWithRedis(token string) (*model.UserSession, error) {
	result := database.GlobalRedis.Get(context.Background(), getKey(token))
	if result.Err() == redis.Nil {
		return nil, errors.New("token not found")
	}
	if result.Err() != nil {
		return nil, result.Err()
	}

	var session model.UserSession
	err := json.Unmarshal([]byte(result.Val()), &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
