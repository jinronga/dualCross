package xdb

import (
	"testing"
)

func TestIndexPolicy_String(t *testing.T) {

	searcher, err := CreateSearcher("../../../db/ip2region.xdb", "vectorIndex")

	if err != nil {
		t.Fatalf("failed to create searcher: %s", err)
	}

	defer func() {
		searcher.Close()
	}()

	str, err := searcher.SearchByStr("162.19.192.82")
	if err != nil {
		t.Fatalf("failed to search: %s", err)
		return
	}
	t.Logf("search result: %s", str)
}
