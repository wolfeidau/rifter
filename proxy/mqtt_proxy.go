package proxy

import (
	"github.com/wolfeidau/rifter/metrics"
	"time"
)

type ProxySettings struct {
	DialTimeout time.Duration
}

type ReverseProxy struct {
	metrics metrics.ProxyMetrics
}
