package graphwriter

//
// import (
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"
//
// 	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
// )
//
// func neo4jWriter(g *GraphWriter) {
// 	driver, err := neo4j.NewDriver(fmt.Sprintf("neo4j://%v", endpoint), neo4j.BasicAuth("neo4j", "test", ""))
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer driver.Close()
// 	fmt.Println("neo4j client connected to", endpoint)
//
// 	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
// 	defer session.Close()
//
// 	batch := []string{}
//
// 	sendData := func() {
// 		fmt.Println("sending", len(batch), "cypher commands")
// 		cypherData := strings.Join(batch, "")
// 		// fmt.Println(cypherData)
//
// 		_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
// 			result, err := transaction.Run(
// 				cypherData,
// 				map[string]interface{}{})
// 			if err != nil {
// 				return nil, err
// 			}
// 			return nil, result.Err()
// 		})
// 		batch = []string{}
// 	}
//
// 	timeout := time.Duration(time.Duration(time.Second))
// 	timer := time.NewTimer(timeout)
// 	for {
// 		select {
// 		case <-timer.C:
// 			if len(batch) > 0 {
// 				sendData()
// 			}
// 			timer.Reset(timeout)
// 		case m := <-cypherChan:
// 			batch = append(batch, m)
//
// 			if len(batch) >= 1000 {
// 				sendData()
// 				if !timer.Stop() {
// 					<-timer.C
// 				}
// 				timer.Reset(timeout)
// 			}
// 		}
// 	}
// }
