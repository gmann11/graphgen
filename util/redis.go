package util

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisSender(operation, cypherQuery, endpoint string) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:6379", endpoint),
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
