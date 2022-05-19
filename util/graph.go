package util

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func GenerateGraph(siteCount, productCount, attributeCount, batchSize int, redis, neo4j bool, workers int) {

	if neo4j {
		fmt.Println("running the neo4j inserts")
		sendSiteNodes(siteCount, Neo4jSender, batchSize)
		sendProductNodes(productCount, attributeCount, Neo4jSender, batchSize, workers)
	}

	if redis {
		fmt.Println("running the redis inserts")
		sendSiteNodes(siteCount, RedisSender, batchSize)
		sendProductNodes(productCount, attributeCount, RedisSender, batchSize, workers)
	}

}

func sendSiteNodes(s int, writer func(string, string), batchSize int) {

	var wg sync.WaitGroup

	cypherChan := make(chan (string))

	go func() {
		for i := 0; i < s; i++ {
			cypherChan <- CreateSite(i)
		}
		close(cypherChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		var counter int
		batch := []string{}
		for cypher := range cypherChan {

			if counter >= batchSize {
				writer("CREATE", sliceToCypher(batch))
				batch = []string{}
				counter = 0
			}
			batch = append(batch, cypher)
			counter += 1
		}
		writer("CREATE", sliceToCypher(batch))
	}()
	wg.Wait()
}

func sendProductNodes(productCount, attributeCount int, writer func(string, string), batchSize int, workers int) {

	var wg sync.WaitGroup

	cypherChan := make(chan (string))

	var eps uint32
	done := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
	Loop:
		for {
			select {
			case <-done:
				break Loop
			case <-ticker.C:
				log.Println("Events per second:", eps)
				eps = 0
			}
		}
	}()

	// Producer
	go func() {
		for i := 0; i < productCount; i++ {
			cypherChan <- CreateProduct(i, attributeCount)
		}
		close(cypherChan)
	}()

	// Consumer
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var batchCounter int
			batch := []string{}
			for cypher := range cypherChan {
				atomic.AddUint32(&eps, 1)

				batch = append(batch, cypher)
				batchCounter += 1

				if batchCounter >= batchSize {
					writer("CREATE", sliceToCypher(batch))
					batch = []string{}
					batchCounter = 0
				}

			}
			// send any remaining products
			if len(batch) > 0 {
				writer("CREATE", sliceToCypher(batch))
			}
		}()
	}

	wg.Wait()
	done <- true
}
