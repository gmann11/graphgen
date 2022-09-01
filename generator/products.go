package generator

import (
	"fmt"
	"graphgen/util"
	"math/rand"

	"github.com/brianvoe/gofakeit"
)

// createProduct generates a CREATE cypher command for a product
// attributes are split 1/3 ints, 1/3 floats, 1/3 strings
func createProduct(id, attributes int) string {
	product := make(map[string]interface{})

	d := attributes / 3
	r := attributes % 3

	// add the guideID
	product["id"] = fmt.Sprintf("%v", id)
	product["countryCode"] = fmt.Sprintf("'%v'", gofakeit.CountryAbr())
	product["longitude"] = gofakeit.Longitude()
	product["latitude"] = gofakeit.Latitude()

	// add 1/3 ints
	for i := 0; i < d; i++ {
		product[fmt.Sprintf("a%d", i)] = fmt.Sprintf("%d", rand.Int31n(100))
	}

	// add 1/3 floats
	for i := (d * 1); i < (d * 2); i++ {
		product[fmt.Sprintf("a%d", i)] = fmt.Sprintf("%.2f", rand.Float32())
	}

	// add 1/3 strings plus any remainders
	for i := (d * 2); i < (d*3)+r+1; i++ {
		product[fmt.Sprintf("a%d", i)] = fmt.Sprintf("'%v'", util.RandomWord())
	}

	return fmt.Sprintf("CREATE (p:product {%v})", util.MapToCypher(product))
}

func linkProductToSite(productID, site int) string {
	edge := make(map[string]interface{})
	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", util.RandomWord())

	return fmt.Sprintf("MATCH (p:product),(s:site) WHERE p.id=%v AND s.id=%v CREATE (p)-[r:cached {%v}]->(s)", productID, site, util.MapToCypher(edge))
}

func createProductEdge(productID, siteCount, productLinkage int) []string {
	var edges []string
	for len(edges) < productLinkage {
		// get a random site
		randomSite := rand.Intn(siteCount)
		edges = append(edges, linkProductToSite(productID, randomSite))
	}
	return edges
}
