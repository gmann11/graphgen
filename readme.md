# graph gen

## how to build

```sh
go mod init graphgen
go mod tidy
go get -u
go run ./main.go
```

## how to get help

```sh
go run ./main.go create -h

❯ go run ./main.go create -h
dictionary loaded with 235887 words
create a graph and send to a backend

Usage:
  graphgen create [flags]

Flags:
  -a, --attributes int         number of attributes (default 40)
  -b, --batch int              batch size (default 100)
  -h, --help                   help for create
      --neo4j                  send to neo4j (default true)
  -n, --neo4jEndpoint string   endpoint for neo4j (default "localhost:7687")
      --productLinkage int     product linkage (default 4)
  -p, --products int           number of products (default 1000)
      --redis                  send to redis (default true)
  -r, --redisEndpoint string   endpoint for redis (default "localhost:6379")
      --siteLinkage int        site linkage (default 4)
  -s, --sites int              number of sites (default 25)
```

## local testing

run neo4j:

```sh
#!/bin/bash

export USID="$(id -u ${USER})" 
export GRID="$(id -g ${USER})"
export BASE="/home/ec2-user/work/jren/run/ce-db"
PUBIP=`curl -s http://169.254.169.254/latest/meta-data/public-hostname`

docker run \
    -p7474:7474 -p7687:7687 \
    -u $(id -u ${USER}):$(id -g ${USER}) \
    -v $BASE/data:/data \
    -v $BASE/logs:/logs \
    -v $BASE/import:/var/lib/neo4j/import \
    -v $BASE/plugins:/plugins \
    --env NEO4J_AUTH=neo4j/password \
    --env NEO4J_dbms_memory_pagecache_size=1G \
    --env NEO4J_dbms_memory_heap_initial__size=512M \
    --env NEO4J_dbms_memory_heap_max__size=512M \
    --env NEO4J_dbms_connector_bolt_advertised__address=${PUBIP}:7687 \
    --env NEO4J_dbms_connector_http_advertised__address=${PUBIP}:7474 \
    --env NEO4J_dbms_security_procedures_unrestricted=apoc.* \
    --env NEO4J_ACCEPT_LICENSE_AGREEMENT=yes \
    --env NEO4J_dbms_logs_query_transaction__id.enabled=true \
    --env NEO4J_dbms_logs_query_enabled=OFF \
    --env NEO4J_dbms_tx__log_rotation_retention__policy=false \
    --name neo4j \
    neo4j:4.4.10-community
```
run redis:

```sh
docker run --rm -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```

queries:
```sh
cd query
go run ./main.go [add]
```
