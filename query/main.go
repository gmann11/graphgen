package main

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"fmt"
	"os"
)

var  queries = []string{
  "CALL apoc.warmup.run(true,true,true)",
  "MATCH (n) return count(n)",
  "MATCH (p:Product) WHERE p.a26='behavior' RETURN count(p)",
  "MATCH (p:Product) WHERE p.country='US' RETURN count(p)",
  "MATCH (s:Site {id:0}),(d:Site {id:19}) RETURN shortestPath((s)-[:CONNECTED*]->(d))",
  `WITH point({longitude: 131.132813, latitude: 25.005979}) AS lowerLeft,
  point({longitude: -59.414063, latitude: 50.28933}) AS upperRight
  MATCH (p:Product) WHERE point.withinBBox(point({longitude: p.longitude, latitude:p.latitude}),lowerLeft, upperRight)
  RETURN count(p)`,
  `MATCH (p:Product) WHERE p.latitude > 25.005979 AND p.latitude < 50.28933
  AND p.longitude > -131.132813 AND p.longitude < -59.414063
  RETURN count(p)`,
  "MATCH (p:Product) RETURN avg(p.a10)",
  `MATCH (p:Product) WHERE point.withinBBox(p.geo,
   point({longitude: 131.132813, latitude: 25.005979}), point({longitude: -59.414063, latitude: 50.280933}))
   RETURN count(p)`,
  }

var  queries_add = []string{
  `WITH point({longitude: 131.132813, latitude: 25.005979}) AS lowerLeft,
  point({longitude: -59.414063, latitude: 50.28933}) AS upperRight
  MATCH (p:product) WHERE point.withinBBox(p.geo,lowerLeft, upperRight)
  RETURN count(p)`,
  `MATCH (p:product) WHERE point.withinBBox(p.geo,
  point({longitude: 131.132813, latitude: 25.005979}), point({longitude: -59.414063, latitude: 50.280933}))
  RETURN count(p)`,
  "MATCH (s:site {id:18})-[:connected*3..6]->(s2:site) RETURN count(s2)",
  "MATCH (s:site {id:18})-[c:connected]->(s2:site)-[c2:connected]-(s3:site) return sum(c2.a2)",
  `MATCH (site:site { id:18 })-[:connected*1..2]-(othersite:site) WHERE site <> othersite
  WITH DISTINCT othersite
  MATCH (othersite)<-[c:cached]-(prod) WHERE prod.a23 > .2
  WITH prod, collect(othersite) AS othersites
  RETURN prod, othersites LIMIT 20`,
}

func query(d neo4j.Driver, q []string) {
  session := d.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead,})
  defer session.Close()
  for _, value := range q {
    session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
      result, err := transaction.Run(value, nil)
      if err != nil {
        return nil, err
      }
     rs,_ := result.Consume()
      fmt.Printf("Query: %s\n", rs.Query().Text())
      fmt.Printf("Returned in: %d ns, consumed in: %d ns\n", rs.ResultAvailableAfter(), rs.ResultConsumedAfter())
      return nil, nil
    })
  }
}

func main() {
  args := os.Args[1:]
  d := driver()
  if len(args) > 0 && args[0] == "add" {
    query(d, queries_add)
  } else {
    query(d, queries)
  }
}


func driver() neo4j.Driver {
	endpoint := "172.31.69.193:7687"
	driver, err := neo4j.NewDriver(fmt.Sprintf("neo4j://%v", endpoint), neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		panic(err)
	}
	return driver
}
