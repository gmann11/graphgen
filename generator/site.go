package generator

import (
	"graphgen/util"
	"math/rand"
)

// createSite generates data for a site
func createSite(id int) map[string]interface{} {
	site := make(map[string]interface{})

	// add the site id
	site["id"] = id

	// add an int
	site["a1"] = rand.Int31n(100)

	// add a float
	site["a2"] = util.RoundFloat(rand.Float64(),2)

	// add a string
	site["a3"] = util.RandomWord()
        return site
}

func linkSiteToSite(ss, ds int) map[string]interface{} {
	edge := make(map[string]interface{})
	edge["from"] = ss
	edge["to"] = ds
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

//func createSiteEdge(id, siteCount, linkage int) []string {
func createSiteEdge(id, siteCount, linkage int) []map[string]interface{} {
	var edges []map[string]interface{}

	for len(edges) < linkage {

		// get a random destination site
		randomSite := rand.Intn(siteCount)

		// if you are that site, skip
		if randomSite == id {
			continue
		}

		edges = append(edges, linkSiteToSite(id, randomSite))
	}
	return edges
}
