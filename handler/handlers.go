package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type DummyHandler struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewDummyHandler(ctx context.Context, redisClient *redis.Client) *DummyHandler {
	return &DummyHandler{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (handler *DummyHandler) GetDummyHandler(c *gin.Context) {
	log.Println("calling DummyHandler...")

	val, err := handler.redisClient.Get(handler.ctx, "redisservice").Result()

	if err == redis.Nil {
		c.JSON(http.StatusOK, gin.H{
			"Message": "first dummy from container...",
		})
		handler.redisClient.Set(handler.ctx, "redisservice", "dummy from container...", 0)
	} else {

		c.JSON(http.StatusOK, gin.H{
			"Message": val,
		})
	}

}
