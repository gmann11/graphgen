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

run redis:

```sh
docker run --rm -p 7474:7474 -p 7687:7687 --env NEO4J_AUTH=neo4j/test neo4j:latest
```

run neo4j:

```sh
docker run --rm -p 6379:6379 -p 8001:8001 redis/redis-stack:latest
```
