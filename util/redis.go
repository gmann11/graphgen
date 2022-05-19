package util

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisSender(operation, cypherQuery string) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := rdb.Do(ctx,
		"GRAPH.QUERY",
		"jren",
		fmt.Sprintf("%v %v", operation, cypherQuery),
	).Err()
	if err != nil {
		log.Println("error execing pipeline", err)
	}
}
