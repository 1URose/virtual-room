package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type Connection struct {
	Client *redis.Client
}

func NewRedisClient(ctx context.Context) *Connection {
	return &Connection{
		Client: NewRedisConnection(ctx).Client,
	}
}

func NewRedisConnection(ctx context.Context) *Connection {
	config := ReadConfigFromEnvironment()

	client := redis.NewClient(&redis.Options{
		Addr: config.Host + ":" + strconv.Itoa(config.Port),
		//Username: config.User,
		//Password: config.Password,
		DB: 0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}
	fmt.Println("Успешно подключились к Redis")

	return &Connection{client}
}
