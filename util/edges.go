package util

import (
	"fmt"
	"math/rand"
)

// MATCH (s:site),(p:product) WHERE s.id=0 AND p.id = 7 CREATE (s)-[:cached]->(p)
func linkProductsToSites(id, site int) string {

	edge := make(map[string]interface{})

	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", randomWord())

	return fmt.Sprintf("MATCH (p:product),(s:site) WHERE p.id=%v AND s.id=%v CREATE (p)-[r:cached {%v}]->(s)", id, site, mapToCypher(edge))

}

func linkSitesToSites(id, site int) string {

	edge := make(map[string]interface{})

	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", randomWord())

	return fmt.Sprintf("MATCH (ss:site),(ds:site) WHERE ss.id=%v AND ds.id=%v CREATE (ss)-[r:connected {%v}]->(ds)", id, site, mapToCypher(edge))

}
