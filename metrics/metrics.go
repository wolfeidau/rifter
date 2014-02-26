package metrics

import (
	gmetrics "github.com/rcrowley/go-metrics"
)

// Record proxy related meters to enable monitoring
// of throughput and volume.
type ProxyMetrics struct {
	Msgs        gmetrics.Meter
	MsgReply    gmetrics.Meter
	MsgForward  gmetrics.Meter
	MsgBodySize gmetrics.Histogram
}

func NewProxyMetrics() ProxyMetrics {

	pm := ProxyMetrics{
		Msgs:        gmetrics.NewMeter(),
		MsgReply:    gmetrics.NewMeter(),
		MsgForward:  gmetrics.NewMeter(),
		MsgBodySize: gmetrics.NewHistogram(gmetrics.NewExpDecaySample(1028, 0.015)),
	}

	gmetrics.Register("rifter.proxy.msgs", pm.Msgs)
	gmetrics.Register("rifter.proxy.msg_reply", pm.Msgs)
	gmetrics.Register("rifter.proxy.msg_forward", pm.Msgs)
	gmetrics.Register("rifter.proxy.msg_body_size", pm.MsgBodySize)

	return pm
}
