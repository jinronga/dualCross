package ip

type (
	// Query interface
	Query interface {
		// Query query
		Query(ip string) (*LocationInfo, error)
		// QueryResponse query response
		QueryResponse(ip []string) ([]*LocationInfo, error)
	}

	// LocationInfo  location info
	LocationInfo struct {
		IP      string `json:"ip"`
		Country string `json:"country"`
		// 省份
		Province string `json:"region"`
		City     string `json:"city"`
		// 运营商
		ISP string `json:"isp"`
		// 城市编码
		Code string `json:"code"`
	}
)
