package ip

import (
	"dualcross/internal/conf"
	"dualcross/pkg/ip2region/xdb"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type (
	// IP2RegionSearcher is a searcher for ip2region.
	IP2RegionSearcher struct {
		searcher *xdb.Searcher
	}
)

// NewIP2RegionSearcher creates a new IP2RegionSearcher.
func NewIP2RegionSearcher(c *conf.Ip2Region) (Query, error) {
	searcher, err := xdb.CreateSearcher(c.GetDbPath(), "vectorIndex")
	if err != nil {
		return nil, err
	}

	return &IP2RegionSearcher{
		searcher: searcher,
	}, nil
}

// LocationInfoPool 对象池减少GC压力
var LocationInfoPool = sync.Pool{
	New: func() interface{} {
		return &LocationInfo{}
	},
}

func (i *IP2RegionSearcher) Query(ip string) (*LocationInfo, error) {
	res, err := i.searcher.SearchByStr(ip)
	if err != nil {
		return nil, err
	}

	return ParseLocationInfo(res)
}

func (i *IP2RegionSearcher) QueryResponse(ips []string) ([]*LocationInfo, error) {

	input := make(chan string, 1000)
	results := make(chan *LocationInfo, 1000)
	go func() {
		for _, ip := range ips {
			r, err := i.searcher.SearchByStr(ip)
			if err != nil {
				continue
			}
			input <- r
		}
		close(input)
	}()
	go BatchProcessor(input, results)
	var resIPs []*LocationInfo
	for info := range results {
		resIPs = append(resIPs, info)
	}
	return resIPs, nil
}

func BatchProcessor(input <-chan string, results chan<- *LocationInfo) {
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU()*2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for s := range input {
				if info, err := ParseLocationInfo(s); err == nil {
					results <- info
				} else {
					fmt.Printf("Failed to parse location info: %q, error: %v\n", s, err)
				}
			}
		}()
	}
	wg.Wait()
	close(results)
}

func ParseLocationInfo(s string) (*LocationInfo, error) {
	parts := strings.Split(s, "|")
	if len(parts) < 5 {
		return nil, fmt.Errorf("invalid formatxxxxxxxxxxxxxxxxxx")
	}

	info := LocationInfoPool.Get().(*LocationInfo)
	info.Country = strings.TrimSpace(parts[0])
	info.Code = StringStrip(parts[1])
	info.Province = strings.TrimSpace(parts[2])
	info.City = strings.TrimSpace(parts[3])
	info.ISP = strings.TrimSpace(parts[4])
	return info, nil
}

func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}
