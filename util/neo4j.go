package util

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func neo4jSender(cypherChan chan (string), endpoint string) {
	driver, err := neo4j.NewDriver(fmt.Sprintf("neo4j://%v:7687", endpoint), neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		log.Panic(err)
	}
	defer driver.Close()
	fmt.Println("neo4j client connected to", endpoint)

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	for cypherQuery := range cypherChan {
		_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				fmt.Sprintf("%v", cypherQuery),
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
}
