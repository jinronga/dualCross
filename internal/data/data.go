package data

import (
	"dualcross/internal/conf"
	"dualcross/pkg/ip"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	IPQuery ip.Query
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	ipQuery, err := newIPQuery(c.GetIpdb())
	if err != nil {
		return nil, nil, err
	}
	return &Data{
		IPQuery: ipQuery,
	}, cleanup, nil
}

func newIPQuery(c *conf.IpDBInfo) (ip.Query, error) {
	switch c.Type {
	case "ip2region":
		return ip.NewIP2RegionSearcher(c.GetRegion())
	case "geoip":

	default:
	}
	return nil, nil
}
