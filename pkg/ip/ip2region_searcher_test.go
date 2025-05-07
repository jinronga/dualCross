package ip

import (
	"dualcross/internal/conf"
	"testing"
)

func TestParseLocationInfoHighPerf(t *testing.T) {
	c := &conf.Ip2Region{DbPath: "../../db/ip2region.xdb"}
	searcher, err := NewIP2RegionSearcher(c)
	if err != nil {
		t.Errorf("NewIP2RegionSearcher err: %v", err)
		return
	}
	ips := []string{
		"162.19.192.82",
	}
	query, err := searcher.QueryResponse(ips)
	if err != nil {
		t.Errorf("Query err: %v", err)
		return
	}
	t.Log(query)
}
