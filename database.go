package main

import(
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var Ctx = context.Background()

func CreateClient() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:		os.Getenv("DB_ADDR"),
		Password:	os.Getenv("DB_PWD"),
		DB: 		0	
	})

	return rdb
}