package infrastructure

import (
	"context"
	"git.ai-space.tech/coursework/backend/internal/infrastructure/cache/redis"
	postgresConnector "git.ai-space.tech/coursework/backend/internal/infrastructure/db/postgres"
	"github.com/jackc/pgx/v5"
	"log"
)

type Connections struct {
	PostgresConnection *pgx.Conn
	RedisClient        *redis.Connection
	Ctx                context.Context
}

func NewConnections(ctx context.Context) *Connections {
	postgres := postgresConnector.NewConnection(ctx).Conn
	redisClient := redis.NewRedisConnection(ctx)

	return &Connections{
		PostgresConnection: postgres,
		RedisClient:        redisClient,
		Ctx:                ctx,
	}
}

func (c *Connections) Close() {
	if c.PostgresConnection != nil {
		err := c.PostgresConnection.Close(c.Ctx)
		if err != nil {
			log.Fatalf("Failed to close PostgreSQL connection: %v\n", err)
		}
	}

	/*if c.RedisClient != nil {
		err := c.RedisClient.Client.Close()
		if err != nil {
			log.Fatalf("Failed to close Redis connection: %v\n", err)
		}
	}*/
}
