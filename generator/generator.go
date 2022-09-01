package generator

import (
	"fmt"
)

type CypherCommands struct {
	SiteNodes    []string
	SiteEdges    []string
	ProductNodes []string
	ProductEdges []string
	Indexes      []string
}

func NewCypherCommands() *CypherCommands {
	fmt.Println("generating cypher commands in memory")

	cc := CypherCommands{
		Indexes: []string{
			"CREATE INDEX FOR (s:site) ON (s.id)",
			"CREATE INDEX FOR (p:product) ON (p.id)",
		},
	}
	return &cc
}

// Generates sites
func (c *CypherCommands) GenerateSiteNodes(siteCount int) {
	fmt.Println("generating site nodes")
	for i := 0; i < siteCount; i++ {
		s := createSite(i)
		c.SiteNodes = append(c.SiteNodes, s)
	}
}

// Generates products
func (c *CypherCommands) GenerateProductNodes(productCount, attributeCount int) {
	fmt.Println("generating product nodes")
	for i := 0; i < productCount; i++ {
		p := createProduct(i, attributeCount)
		c.ProductNodes = append(c.ProductNodes, p)
	}
}

// GenerateSiteEdges links sites to other sites
func (c *CypherCommands) GenerateSiteEdges(siteCount, siteLinkage int) {
	fmt.Println("generating site-to-site edges")
	for i := 0; i < siteCount; i++ {
		for _, edge := range createSiteEdge(i, siteCount, siteLinkage) {
			c.SiteEdges = append(c.SiteEdges, edge)
		}
	}
}

// GenerateProductEdges links products to sites
func (c *CypherCommands) GenerateProductEdges(siteCount, productCount, productLinkage int) {
	fmt.Println("generating product-to-site edges")
	for i := 0; i < productCount; i++ {
		for _, edge := range createProductEdge(i, siteCount, productLinkage) {
			c.ProductEdges = append(c.ProductEdges, edge)
		}
	}
}
