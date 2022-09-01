package cmd

import (
	"fmt"
	"graphgen/generator"
	"graphgen/graphwriter"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a graph",
	Long:  "create a graph and send to a backend",
	Run: func(cmd *cobra.Command, args []string) {
		sites, _ := cmd.Flags().GetInt("sites")
		siteLinkage, _ := cmd.Flags().GetInt("siteLinkage")
		productLinkage, _ := cmd.Flags().GetInt("productLinkage")
		products, _ := cmd.Flags().GetInt("products")
		attributes, _ := cmd.Flags().GetInt("attributes")
		batchSize, _ := cmd.Flags().GetInt("batch")
		redis, _ := cmd.Flags().GetBool("redis")
		neo4j, _ := cmd.Flags().GetBool("neo4j")
		redisep, _ := cmd.Flags().GetString("redisEndpoint")
		neo4jep, _ := cmd.Flags().GetString("neo4jEndpoint")

		fmt.Println("generating a graph with the following parameters")
		fmt.Println("sites", sites)
		fmt.Println("site linkage", siteLinkage)
		fmt.Println("products", products)
		fmt.Println("product linkage", productLinkage)
		fmt.Println("attributes", attributes)
		fmt.Println("batch size", batchSize)
		fmt.Println("redis", redis)
		fmt.Println("neo4j", neo4j)
		fmt.Println("redisEndpoint", redisep)
		fmt.Println("neo4jEndpoint", neo4jep)

		if siteLinkage > sites {
			log.Panic("can create more site links than there are sites")
		}
		if productLinkage > sites {
			log.Panic("can not link products to more sites than there are sites")
		}

		cc := generator.NewCypherCommands()
		cc.GenerateSiteNodes(sites)
		cc.GenerateSiteEdges(sites, siteLinkage)
		cc.GenerateProductNodes(products, attributes)
		cc.GenerateProductEdges(sites, products, productLinkage)

		graphdbs := []*graphwriter.GraphWriter{}

		if neo4j {
			gw := graphwriter.NewGraphWriter(graphwriter.Neo4j, batchSize, neo4jep)
			graphdbs = append(graphdbs, gw)
		}
		if redis {
			gw := graphwriter.NewGraphWriter(graphwriter.Redis, batchSize, redisep)
			graphdbs = append(graphdbs, gw)
		}

		for _, gw := range graphdbs {

			start := time.Now()

			log.Println("running test for", gw.Name)
			// add site nodes
			log.Println("inserting indexes")
			gw.Write(cc.Indexes)
			log.Println("inserting site nodes")
			gw.Write(cc.SiteNodes)
			log.Println("inserting product nodes")
			gw.Write(cc.ProductNodes)
			log.Println("inserting site edges")
			gw.Write(cc.SiteEdges)
			log.Println("inserting product edges")
			gw.Write(cc.ProductEdges)
			// finally
			results := gw.Close()
			duration := time.Since(start)
			eps := float64(results) / float64(duration.Seconds())
			log.Println("sent", results, "in", time.Since(start), "eps", eps)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntP("sites", "s", 25, "number of sites")
	createCmd.Flags().Int("siteLinkage", 4, "site linkage")
	createCmd.Flags().Int("productLinkage", 4, "product linkage")
	createCmd.Flags().IntP("products", "p", 1000, "number of products")
	createCmd.Flags().IntP("attributes", "a", 40, "number of attributes")
	createCmd.Flags().IntP("batch", "b", 100, "batch size")
	createCmd.Flags().Bool("redis", true, "send to redis")
	createCmd.Flags().Bool("neo4j", true, "send to neo4j")
	createCmd.Flags().StringP("redisEndpoint", "r", "localhost:6379", "endpoint for redis")
	createCmd.Flags().StringP("neo4jEndpoint", "n", "localhost:7687", "endpoint for neo4j")
}
