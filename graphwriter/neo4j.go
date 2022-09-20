package graphwriter

import (
	"fmt"
	"log"
        "strings"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func neo4jWriter(g *GraphWriter) {
	defer g.wg.Done()
	log.Println("connecting to neo4j", g.endpoint)
	driver, err := neo4j.NewDriver(fmt.Sprintf("neo4j://%v", g.endpoint), neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		log.Panic(err)
	}
	defer driver.Close()
	log.Println("neo4j client connected to", g.endpoint)

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	batch := []map[string]interface{}{}

	// iterates through a batch of commands, and writes the transactions
	sendData := func(cyp string) {
		//fmt.Println("sending", len(batch))
		// cypherData := strings.Join(batch, ";\n")
		// // cypher commands separated by ; do not work like in the browser UI
		// fmt.Print(cypherData)
	        _, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(cyp,map[string]interface{}{"dict": batch})
				if err != nil {
					return nil, err
				}
				return result.Consume()
			})

		batch = []map[string]interface{}{}
	}

	var curcyp string
	for message := range g.ch {
		if val, ok := message["cypher"]; ok {
		  if !strings.Contains(val.(string), "$") {
		    sendData(val.(string))
	          } else {
                    if curcyp !=  "" {
	              sendData(curcyp)
		    }
		    curcyp = val.(string)
	          }
	        } else {
		  batch = append(batch, message)
		  if len(batch) >= g.batchSize {
			sendData(curcyp)
		  }
		}
	}
	// final flush
	sendData(curcyp)
}
