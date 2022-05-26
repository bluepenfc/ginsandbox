package main

import (
	"context"
	"log"

	"fc.com/dockernoexternalservice/handler"
	//"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var redisClient *redis.Client
var dhandler *handler.DummyHandler

func init() {
	ctx = context.Background()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping(ctx)
	log.Println(status)
	dhandler = handler.NewDummyHandler(ctx, redisClient)
}

func main() {
	router := gin.Default()

	router.GET("/", dhandler.GetDummyHandler)

	router.Run(":8090")
}
