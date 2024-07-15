package main

import (
	"fmt"
	"time"
)

func main() {
	processor := NewDocumentProcessor()

	doc1 := &Document{
		Uri:       "http://example.com/doc1",
		PubDate:   uint64(time.Now().Unix()),
		FetchTime: uint64(time.Now().Unix()),
		Text:      "First version",
	}

	doc2 := &Document{
		Uri:       "http://example.com/doc1",
		PubDate:   uint64(time.Now().Unix() + 10),
		FetchTime: uint64(time.Now().Unix() + 10),
		Text:      "Second version",
	}

	_, err := processor.Process(doc1)
	if err != nil {
		fmt.Println("Error processing document:", err)
		return
	}

	_, err = processor.Process(doc2)
	if err != nil {
		fmt.Println("Error processing document:", err)
		return
	}

	for _, doc := range processor.docMap["http://example.com/doc1"] {
		fmt.Printf("Uri: %s, PubDate: %d, FetchTime: %d, Text: %s, FirstFetchTime: %d\n",
			doc.Uri, doc.PubDate, doc.FetchTime, doc.Text, doc.FirstFetchTime)
	}
}
