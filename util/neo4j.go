package util

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Neo4jSender(operation, cypherQuery string) {
	driver, err := neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		log.Panic(err)
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			fmt.Sprintf("%v %v", operation, cypherQuery),
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		// log.Println(result)

		return nil, result.Err()
	})
	if err != nil {
		log.Panic(err)
	}
}
