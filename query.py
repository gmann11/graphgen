import sys
import os.path
from neo4j import GraphDatabase
import time 
import datetime

uri = "neo4j://localhost:7687"
driver = GraphDatabase.driver(uri, auth=("neo4j", "password"))

queries = [
  "CALL apoc.warmup.run(true,true,true)",
  "MATCH (n) return count(n)",
  "MATCH (p:product) WHERE p.a26='herpes' RETURN count(p)",
  "MATCH (p:product) WHERE p.country='US' RETURN count(p)",
  "MATCH (s:site {id:0}),(d:site {id:19}) RETURN shortestPath((s)-[:connected*]->(d))",
  """
  WITH point({longitude: 131.132813, latitude: 25.005979}) AS lowerLeft,
  point({longitude: -59.414063, latitude: 50.28933}) AS upperRight
  MATCH (p:product) WHERE point.withinBBox(point({longitude: p.longitude, latitude:p.latitude}),lowerLeft, upperRight)
  RETURN count(p)
  """,
  """
  MATCH (p:product) WHERE p.latitude > 25.005979 AND p.latitude < 50.28933
  AND p.longitude > -131.132813 AND p.longitude < -59.414063
  RETURN count(p)
  """,
  "MATCH (p:product) RETURN avg(p.a10)",
  ]

queries_additional = [
  """
  WITH point({longitude: 131.132813, latitude: 25.005979}) AS lowerLeft,
  point({longitude: -59.414063, latitude: 50.28933}) AS upperRight
  MATCH (p:product) WHERE point.withinBBox(p.geo,lowerLeft, upperRight)
  RETURN count(p)
  """,
  """
  MATCH (p:product) WHERE point.withinBBox(p.geo, 
  point({longitude: 131.132813, latitude: 25.005979}), point({longitude: -59.414063, latitude: 50.280933}))
  RETURN count(p)
  """,
  "MATCH (s:site {id:18})-[:connected*3..6]->(s2:site) RETURN count(s2)",
  "MATCH (s:site {id:18})-[c:connected]->(s2:site)-[c2:connected]-(s3:site) return sum(c2.a2)",
  """
  MATCH (site:site { id:18 })-[:connected*1..2]-(othersite:site) WHERE site <> othersite
  WITH DISTINCT othersite
  MATCH (othersite)<-[c:cached]-(prod) WHERE prod.a23 > .2
  WITH prod, collect(othersite) AS othersites
  RETURN prod, othersites LIMIT 20
"""
]

def runAddQueries():
  with driver.session() as session:
    for q in queries_additional:
      rs = session.run(q).consume()
      print(f"Query: {rs.query}\nReturned in: {rs.result_available_after} ms, Consumed in: {rs.result_consumed_after} ms")

def runQueries():
  with driver.session() as session:
    for q in queries:
      rs = session.run(q).consume()
      #res = [dict(i) for i in r]
      #rs = r.consume()
      #print(f"Query: {rs.query} - result {res}")
      print(f"Query: {rs.query}\nReturned in: {rs.result_available_after} ms, Consumed in: {rs.result_consumed_after} ms")
  
def main():
  if len(sys.argv) > 1 and sys.argv[1] == "add":
    runAddQueries()
  else:
    runQueries()

if __name__ == "__main__":
  main()

