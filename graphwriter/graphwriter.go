package graphwriter

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type WriterType uint8

const (
	Redis = WriterType(iota)
	Neo4j
)

// Writer returns a *GraphWriter
type GraphWriter struct {
	ch           chan string
	wg           sync.WaitGroup
	endpoint     string
	batchSize    int
	commandsSent uint32
	Name         WriterType
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
		ch:        make(chan string),
		batchSize: batchSize,
		endpoint:  endpoint,
		Name:      writer,
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

func (g *GraphWriter) Write(messages []string) {
	for _, message := range messages {
		atomic.AddUint32(&g.commandsSent, 1)
		g.ch <- message
	}
}
