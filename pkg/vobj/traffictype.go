package vobj

// TrafficType  流量分类枚举
//
//go:generate go run ../../cmd/stringer/cmd.go -type=TrafficType -linecomment
type TrafficType int

const (
	TrafficUnknown          TrafficType = iota
	TrafficLocalOutbound                // 本网出省
	TrafficExternalOutbound             // 异网出省
	TrafficExternal                     // 异网流量
)
