package main

import (
	"sort"
	"sync"
)

type Document struct {
	Uri            string
	PubDate        uint64
	FetchTime      uint64
	Text           string
	FirstFetchTime uint64
}

type Processor interface {
	Process(d *Document) (*Document, error)
}

type DocumentProcessor struct {
	mu     sync.Mutex
	docMap map[string][]*Document
}

func NewDocumentProcessor() *DocumentProcessor {
	return &DocumentProcessor{
		docMap: make(map[string][]*Document),
	}
}

func (dp *DocumentProcessor) Process(d *Document) (*Document, error) {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	docs, exists := dp.docMap[d.Uri]
	if !exists {
		d.FirstFetchTime = d.FetchTime
		dp.docMap[d.Uri] = []*Document{d}
		return d, nil
	}

	docs = append(docs, d)
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].FetchTime < docs[j].FetchTime
	})

	firstFetchTime := docs[0].FetchTime
	for _, doc := range docs {
		if doc.FirstFetchTime == 0 {
			doc.FirstFetchTime = firstFetchTime
		}
	}

	dp.docMap[d.Uri] = docs
	return docs[len(docs)-1], nil
}
