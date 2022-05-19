package util

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func GenerateGraph(siteCount, productCount, attributeCount, batchSize int, redis, neo4j bool, workers int, redisEndpoint, neo4jEndpoint string) {

	if neo4j {
		fmt.Println("running the neo4j inserts")
		sendSiteNodes(siteCount, batchSize, Neo4jSender, neo4jEndpoint)
		sendProductNodes(productCount, attributeCount, batchSize, workers, Neo4jSender, neo4jEndpoint)
		sendProductEdges(productCount, siteCount, Neo4jSender, neo4jEndpoint)
		sendSiteEdges(siteCount, Neo4jSender, neo4jEndpoint)
	}

	if redis {
		fmt.Println("running the redis inserts")
		sendSiteNodes(siteCount, batchSize, RedisSender, redisEndpoint)
		sendProductNodes(productCount, attributeCount, batchSize, workers, RedisSender, redisEndpoint)
		sendProductEdges(productCount, siteCount, RedisSender, redisEndpoint)
		sendSiteEdges(siteCount, RedisSender, redisEndpoint)
	}

}
func sendProductEdges(productCount, siteCount int, writer func(string, string, string), endpoint string) {
	for i := 0; i < productCount; i++ {
		for _, s := range rand.Perm(siteCount)[:rand.Intn(4)+1] {
			cypherQuery := linkProductsToSites(i, s)
			writer("MATCH", cypherQuery, endpoint)
		}
	}
}

func sendSiteEdges(siteCount int, writer func(string, string, string), endpoint string) {
	for i := 0; i < siteCount; i++ {

		sites := makeRange(0, siteCount-1)
		sites = remove(sites, i)
		rand.Shuffle(len(sites), func(i, j int) { sites[i], sites[j] = sites[j], sites[i] })

		for _, s := range sites[:rand.Intn(4)+1] {
			cypherQuery := linkSitesToSites(i, s)
			writer("MATCH", cypherQuery, endpoint)
		}
	}
}

func sendSiteNodes(s int, batchSize int, writer func(string, string, string), endpoint string) {

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
				writer("CREATE", sliceToCypher(batch), endpoint)
				batch = []string{}
				counter = 0
			}
			batch = append(batch, cypher)
			counter += 1
		}
		writer("CREATE", sliceToCypher(batch), endpoint)
	}()
	wg.Wait()
}

func sendProductNodes(productCount, attributeCount, batchSize, workers int, writer func(string, string, string), endpoint string) {

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
					writer("CREATE", sliceToCypher(batch), endpoint)
					batch = []string{}
					batchCounter = 0
				}

			}
			// send any remaining products
			if len(batch) > 0 {
				writer("CREATE", sliceToCypher(batch), endpoint)
			}
		}()
	}

	wg.Wait()
	done <- true
}
