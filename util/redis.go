package util

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func redisSender(cypherChan chan (string), endpoint string) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:6379", endpoint),
	})

	// test redis connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("could not connect to redis streams endpoint", err)
	}
	fmt.Println("redis client connected to", endpoint)

	// send cypher queries send via channel
	for cypherQuery := range cypherChan {
		err := rdb.Do(ctx,
			"GRAPH.QUERY",
			"jren",
			fmt.Sprintf("%v", cypherQuery),
		).Err()
		if err != nil {
			log.Println("redis error executing pipeline", err)
		}
	}
}
