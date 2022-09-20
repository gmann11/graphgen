package generator

import (
	"fmt"
	"graphgen/util"
	"math/rand"

	"github.com/brianvoe/gofakeit"
)

// createProduct generates a CREATE cypher command for a product
// attributes are split 1/3 ints, 1/3 floats, 1/3 strings
func createProduct(id, attributes int) map[string]interface{} {
	product := make(map[string]interface{})

	d := attributes / 3
	r := attributes % 3

	// add the guideID
	product["id"] = id
	product["countryCode"] = gofakeit.CountryAbr()
	product["longitude"] = gofakeit.Longitude()
	product["latitude"] = gofakeit.Latitude()

	// add 1/3 ints
	for i := 0; i < d; i++ {
		product[fmt.Sprintf("a%d", i)] = rand.Int31n(100)
	}

	// add 1/3 floats
	for i := (d * 1); i < (d * 2); i++ {
		product[fmt.Sprintf("a%d", i)] = util.RoundFloat(rand.Float64(),2)
	}

	// add 1/3 strings plus any remainders
	for i := (d * 2); i < (d*3)+r+1; i++ {
		product[fmt.Sprintf("a%d", i)] = util.RandomWord()
	}

	return product
}

func linkProductToSite(productID, site int) map[string]interface{} {
	edge := make(map[string]interface{})
	edge["from"] = productID
	edge["to"] = site
        var atts = make(map[string]interface{})
	// add an int
        atts["a1"] = rand.Int31n(100)

	// add a float
        atts["a2"] = util.RoundFloat(rand.Float64(),2)

	// add a string
        atts["a3"] = util.RandomWord()
	edge["atts"] = atts

	return edge
}

func createProductEdge(productID, siteCount, productLinkage int) []map[string]interface{} {
	var edges []map[string]interface{}
	for len(edges) < productLinkage {
		// get a random site
		randomSite := rand.Intn(siteCount)
		edges = append(edges, linkProductToSite(productID, randomSite))
	}
	return edges
}
