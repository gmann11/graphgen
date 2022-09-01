package graphwriter

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func redisWriter(g *GraphWriter) {
	defer g.wg.Done()
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v", g.endpoint),
	})
	_, err := rdb.FlushAll(ctx).Result()
	if err != nil {
		log.Println("error flushing redis", err)
	}
	pipe := rdb.Pipeline()

	// test redis connection
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("could not connect to redis streams endpoint", err)
	}
	fmt.Println("redis client connected to", g.endpoint)

	sendData := func() {
		fmt.Println("sending", pipe.Len(), "cypher commands")

		_, err := pipe.Exec(ctx)
		if err != nil {
			log.Println("redis error executing pipeline", err)
		}
	}

	for message := range g.ch {
		pipe.Do(ctx, "GRAPH.QUERY", "jren", message)
		if pipe.Len() >= g.batchSize {
			sendData()
		}
	}
	// flush final partial batch
	sendData()
}
