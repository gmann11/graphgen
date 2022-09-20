import sys
import os.path
from neo4j import GraphDatabase
import time 
import datetime

#uri = "neo4j://localhost:7687"
uri = "neo4j://172.31.69.193:7687"
driver = GraphDatabase.driver(uri, auth=("neo4j", "password"))

queries = [
  "CALL apoc.warmup.run(true,true,true)",
  "MATCH (n) return count(n)",
  "MATCH (p:Product) WHERE p.a26='behavior' RETURN count(p)",
  "MATCH (p:Product) WHERE p.country='US' RETURN count(p)",
  "MATCH (s:Site {id:0}),(d:Site {id:19}) RETURN shortestPath((s)-[:CONNECTED*]->(d))",
  """
  WITH point({longitude: 131.132813, latitude: 25.005979}) AS lowerLeft,
  point({longitude: -59.414063, latitude: 50.28933}) AS upperRight
  MATCH (p:Product) WHERE point.withinBBox(point({longitude: p.longitude, latitude:p.latitude}),lowerLeft, upperRight)
  RETURN count(p)
  """,
  """
  MATCH (p:Product) WHERE p.latitude > 25.005979 AND p.latitude < 50.28933
  AND p.longitude > -131.132813 AND p.longitude < -59.414063
  RETURN count(p)
  """,
  "MATCH (p:Product) RETURN avg(p.a10)",
  """
      MATCH (p:Product) WHERE point.withinBBox(p.geo, 
      point({longitude: 131.132813, latitude: 25.005979}), point({longitude: -59.414063, latitude: 50.280933}))
      RETURN count(p)
  """
  ]


def runQueries():
  with driver.session() as session:
    for q in queries:
      r = session.run(q)
      res = [dict(i) for i in r]
      rs = r.consume()
      print(f"Query: {rs.query} - result {res}")
      print(f"Returned in: {rs.result_available_after} ms, consumed in: {rs.result_consumed_after} ms")
  
def main():
  runQueries()

if __name__ == "__main__":
  main()

