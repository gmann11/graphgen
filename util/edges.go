package util

import (
	"fmt"
	"math/rand"
)

// MATCH (s1:site) WHERE s1.id=33
func createSiteMatch(siteAliasID, siteID int) string {
	return fmt.Sprintf("MATCH (s%v:site) WHERE s%v.id=%v", siteAliasID, siteAliasID, siteID)
}

// uses aliases for sites
// CREATE (s1)-[:cached]->(p)
func createProductToSiteEdge(siteAliasID int) string {

	edge := make(map[string]interface{})

	// add an int
	edge["a1"] = fmt.Sprintf("%d", rand.Int31n(100))

	// add a float
	edge["a2"] = fmt.Sprintf("%.2f", rand.Float32())

	// add a string
	edge["a3"] = fmt.Sprintf("'%v'", randomWord())

	return fmt.Sprintf("CREATE (s%v)-[:cached {%v}]->(p)", siteAliasID, mapToCypher(edge))

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
