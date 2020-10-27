package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {
	ch := All(context.Background(), "READMEmd", []string{"../../docs/doc1.txt"})
	results, ok := <-ch
	if !ok {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}
