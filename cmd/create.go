package cmd

import (
	"fmt"
	"graphgen/util"
	"log"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a graph",
	Long:  "create a graph and send to a backend",
	Run: func(cmd *cobra.Command, args []string) {

		sites, _ := cmd.Flags().GetInt("sites")
		products, _ := cmd.Flags().GetInt("products")
		attributes, _ := cmd.Flags().GetInt("attributes")
		batchSize, _ := cmd.Flags().GetInt("batch")
		redis, _ := cmd.Flags().GetBool("redis")
		neo4j, _ := cmd.Flags().GetBool("neo4j")
		workers, _ := cmd.Flags().GetInt("workers")
		redisep, _ := cmd.Flags().GetString("redisEndpoint")
		neo4jep, _ := cmd.Flags().GetString("neo4jEndpoint")

		log.Println("creating a graph with the following parameters")
		fmt.Println("sites", sites)
		fmt.Println("products", products)
		fmt.Println("attributes", attributes)
		fmt.Println("batch size", batchSize)
		fmt.Println("workers", workers)
		fmt.Println("redis", redis)
		fmt.Println("neo4j", neo4j)
		fmt.Println("redisEndpoint", redisep)
		fmt.Println("neo4jEndpoint", neo4jep)

		util.GenerateGraph(cmd)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntP("sites", "s", 25, "number of sites")
	createCmd.Flags().IntP("products", "p", 10_000, "number of products")
	createCmd.Flags().IntP("attributes", "a", 40, "number of attributes")
	createCmd.Flags().IntP("batch", "b", 1000, "batch size")
	createCmd.Flags().Bool("redis", true, "send to redis")
	createCmd.Flags().Bool("neo4j", true, "send to neo4j")
	createCmd.Flags().IntP("workers", "w", 4, "size of worker pool")
	createCmd.Flags().StringP("redisEndpoint", "r", "localhost", "endpoint for redis")
	createCmd.Flags().StringP("neo4jEndpoint", "n", "localhost", "endpoint for neo4j")

}
