package generator

import (
	"fmt"
	"graphgen/util"
	"math/rand"
)

// createSite generates a CREATE cypher command for a site
func createSite(id int) string {
	site := make(map[string]interface{})

	// add the site id
	site["id"] = fmt.Sprintf("%v", id)

	// add an int
	site["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	site["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	site["a3"] = fmt.Sprintf("'%v'", util.RandomWord())

	return fmt.Sprintf("CREATE (:site {%v})", util.MapToCypher(site))
}

func linkSiteToSite(ss, ds int) string {
	edge := make(map[string]interface{})
	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", util.RandomWord())

	return fmt.Sprintf("MATCH (ss:site),(ds:site) WHERE ss.id=%v AND ds.id=%v CREATE (ss)-[r:connected {%v}]->(ds)", ss, ds, util.MapToCypher(edge))
}

func createSiteEdge(id, siteCount, linkage int) []string {
	var edges []string

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
