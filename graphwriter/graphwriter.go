package graphwriter

import (
	"sync"
)

const (
	Redis WriterType = iota
	Neo4j
)

type WriterType int

// Writer returns a *GraphWriter
type GraphWriter struct {
	ch        chan string
	wg        sync.WaitGroup
	endpoint  string
	batchSize int
}

// NewGraphWriter returns a GraphWriter
func NewGraphWriter(writer WriterType, batchSize int, endpoint string) *GraphWriter {
	// create the channel
	g := GraphWriter{
		ch:        make(chan string),
		batchSize: batchSize,
		endpoint:  endpoint,
	}

	// add the waitgroup
	g.wg.Add(1)

	// switch between redis and neo4j
	switch writer {
	case Redis:
		go redisWriter(&g)
		// case Neo4j:
		//   go neo4jWriter(&g)
	}

	return &g
}

func (w *GraphWriter) Close() {
	close(w.ch)
	w.wg.Wait()
}

func (w *GraphWriter) Write(messages []string) {
	for _, message := range messages {
		w.ch <- message
	}
}
