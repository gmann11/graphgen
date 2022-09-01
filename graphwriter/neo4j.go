package graphwriter

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func neo4jWriter(g *GraphWriter) {
	defer g.wg.Done()
	log.Println("connecting to neo4j", g.endpoint)
	driver, err := neo4j.NewDriver(fmt.Sprintf("neo4j://%v", g.endpoint), neo4j.BasicAuth("neo4j", "test", ""))
	if err != nil {
		log.Panic(err)
	}
	defer driver.Close()
	log.Println("neo4j client connected to", g.endpoint)

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	batch := []string{}

	// iterates through a batch of commands, and writes the transactions
	sendData := func() {
		fmt.Println("sending", len(batch), "cypher commands")

		// cypherData := strings.Join(batch, ";\n")
		// // cypher commands separated by ; do not work like in the browser UI
		// fmt.Print(cypherData)

		for _, command := range batch {
			_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				result, err := transaction.Run(
					command,
					map[string]interface{}{})
				if err != nil {
					return nil, err
				}
				return nil, result.Err()
			})
		}

		batch = []string{}
	}

	for message := range g.ch {
		batch = append(batch, message)
		if len(batch) >= g.batchSize {
			sendData()
		}
	}
	// final flush
	sendData()
}
