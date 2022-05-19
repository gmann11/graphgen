package util

import (
	"fmt"
	"math/rand"
)

// GRAPH.QUERY movies "MATCH (s:site),(p:product) WHERE s.siteID = 1 AND p.productID = 1 CREATE (s)-[r:cached {role:'Luke Skywalker'}]->(p) RETURN r"

func CreateProductEdge(n, siteCount, productCount int) string {

	edge := make(map[string]interface{})

	// add the site id
	edge["id"] = fmt.Sprintf("%v", n)

	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", randomWord())

	// return fmt.Sprintf("(:site {%v}),", mapToCypher(edge))

	return "(p:product),(s:site) WHERE p.id=1 AND s.id=1 CREATE (p)-[r:cached {}]->(s)"

}
