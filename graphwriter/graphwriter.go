package graphwriter

import (
	"strconv"
	"sync"
	"sync/atomic"
	"strings"
)

type WriterType uint8

const (
	Redis = WriterType(iota)
	Neo4j
)

// Writer returns a *GraphWriter
type GraphWriter struct {
	ch           chan map[string]interface{}
	wg           sync.WaitGroup
	endpoint     string
	batchSize    int
	commandsSent uint32
	Name         WriterType
	cypher   string
}

func (w WriterType) String() string {
	name := []string{"redis", "neo4j"}
	i := uint8(w)
	switch {
	case i <= uint8(Neo4j):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

// NewGraphWriter returns a GraphWriter
func NewGraphWriter(writer WriterType, batchSize int, endpoint string) *GraphWriter {
	// create the channel
	g := GraphWriter{
		ch:       make(chan map[string]interface{}),
		batchSize: batchSize,
		endpoint:  endpoint,
		Name:      writer,
		cypher: "",
	}

	// add the waitgroup
	g.wg.Add(1)

	// switch between redis and neo4j
	switch writer {
	case Redis:
		go redisWriter(&g)
	case Neo4j:
		go neo4jWriter(&g)
	}

	return &g
}

func (g *GraphWriter) Close() uint32 {
	close(g.ch)
	g.wg.Wait()
	return g.commandsSent
}

func (g *GraphWriter) Write(cypher string, data []map[string]interface{}) {
	if strings.Contains(cypher, "$") {
	  g.ch <- map[string]interface{}{"cypher": cypher}
        } 
	for _, row := range data {
		atomic.AddUint32(&g.commandsSent, 1)
		g.ch <- row
        }
}
