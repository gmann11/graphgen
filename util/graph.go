package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func GenerateGraph(cmd *cobra.Command) {

	redis, _ := cmd.Flags().GetBool("redis")
	neo4j, _ := cmd.Flags().GetBool("neo4j")

	if redis {
		redisEndpoint, _ := cmd.Flags().GetString("redisEndpoint")
		insertData(cmd, "redis", redisSender, redisEndpoint)
	}
	if neo4j {
		neo4jEndpoint, _ := cmd.Flags().GetString("neo4jEndpoint")
		insertData(cmd, "neo4j", neo4jSender, neo4jEndpoint)
	}

}

func insertData(cmd *cobra.Command, name string, sender func(chan (string), string), endpoint string) {
	workers, _ := cmd.Flags().GetInt("workers")

	jobStart := time.Now()
	fmt.Printf("running the %v inserts\n", name)

	cypherChan := make(chan (string))
	// fire up X amount of redis senders
	for i := 0; i < workers; i++ {

		// will close when channel closes
		go sender(cypherChan, endpoint)

	}

	time.Sleep(time.Second)
	// site nodes
	siteCount, _ := cmd.Flags().GetInt("sites")
	fmt.Printf("%v: creating site nodes\n", name)
	start := time.Now()
	for i := 0; i < siteCount; i++ {
		cypherChan <- createSite(i)
	}
	fmt.Printf("%v: %v nodes inserted in %v\n", name, siteCount, time.Since(start))

	// product nodes
	productCount, _ := cmd.Flags().GetInt("products")
	attributeCount, _ := cmd.Flags().GetInt("attributes")
	fmt.Printf("%v: creating product nodes\n", name)
	start = time.Now()
	for i := 0; i < productCount; i++ {
		cypherChan <- createProduct(i, attributeCount)
	}
	duration := time.Since(start)
	eps := float64(productCount) / duration.Seconds()
	fmt.Printf("%v: %v nodes inserted in %v %.2f\n", name, productCount, duration, eps)

	// site edges
	fmt.Printf("%v: linking sites to sites\n", name)
	start = time.Now()
	counter := 0
	for i := 0; i < siteCount; i++ {

		// create a sequence of ints [0,1,2,3...]
		sites := makeRange(0, siteCount-1)
		// remove self
		sites = remove(sites, i)
		//shuffle the output
		rand.Shuffle(len(sites), func(i, j int) { sites[i], sites[j] = sites[j], sites[i] })

		for _, s := range sites[:rand.Intn(4)+1] { //TODO hard-coded 4
			cypherChan <- linkSitesToSites(i, s)
			counter += 1
		}
	}
	duration = time.Since(start)
	eps = float64(counter) / duration.Seconds()
	fmt.Printf("%v: %v edges inserted in %v %.2f\n", name, counter, duration, eps)

	// product edges
	fmt.Printf("%v: linking products to sites\n", name)
	start = time.Now()
	counter = 0
	for i := 0; i < productCount; i++ {
		for _, s := range rand.Perm(siteCount)[:rand.Intn(4)+1] {
			cypherChan <- linkProductsToSites(i, s)
			counter += 1
		}
	}
	duration = time.Since(start)
	eps = float64(counter) / duration.Seconds()
	fmt.Printf("%v: %v edges inserted in %v %.2f\n", name, counter, duration, eps)

	fmt.Printf("%v: all nodes and edges inserted in %v\n\n", name, time.Since(jobStart))

}
