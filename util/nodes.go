package util

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

// CREATE (:site {id:0,a1:56,a2:0.30,a3:'jumpable'})

func createSite(n int) string {

	site := make(map[string]interface{})

	// add the site id
	site["id"] = fmt.Sprintf("%v", n)

	// add an int
	site["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	site["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	site["a3"] = fmt.Sprintf("'%v'", randomWord())

	return fmt.Sprintf("CREATE (:site {%v})", mapToCypher(site))
}

func createProduct(id, attributes int) string {

	// subtract 1 for the GUIDE ID
	// any additional remaining attributes will be strings
	d := (attributes - 1) / 3
	r := (attributes - 1) % 3

	product := make(map[string]interface{})

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
		product[fmt.Sprintf("a%d", i)] = fmt.Sprintf("'%v'", randomWord())
	}

	return fmt.Sprintf("CREATE (p:product {%v})", mapToCypher(product))
}
