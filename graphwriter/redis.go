package graphwriter

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func redisWriter(g *GraphWriter) {
	defer g.wg.Done()
	log.Println("connecting to redis", g.endpoint)
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v", g.endpoint),
	})

	// test redis connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("could not connect to redis streams endpoint", err)
	}
	log.Println("redis client connected to", g.endpoint)

	// flush redis db
	_, err = rdb.FlushAll(ctx).Result()
	if err != nil {
		log.Println("error flushing redis", err)
	}

	// create redis pipeline
	pipe := rdb.Pipeline()

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
