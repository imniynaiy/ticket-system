package database

import (
	"fmt"

	"github.com/imniynaiy/ticket-system/internal/config"
	"github.com/redis/go-redis/v9"
)

var GlobalRedis *redis.Client

func InitRedis(rc *config.RedisConfig) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rc.Address, rc.Port),
		Password: rc.Password, // no password set
		DB:       rc.DB,       // use default DB
	})
	GlobalRedis = rdb
}
