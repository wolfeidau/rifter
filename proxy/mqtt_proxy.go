package proxy

import (
	"github.com/mailgun/vulcan/loadbalance"
	"github.com/mailgun/vulcan/ratelimit"
	"github.com/wolfeidau/rifter/metrics"
	"time"
)

type ProxySettings struct {

	// MemoryBackend or CassandraBackend
	ThrottlerBackend backend.Backend

	// Load balancing algo, e.g. RandomLoadBalancer
	LoadBalancer loadbalance.Balancer

	// How long would proxy wait for server response
	ReadTimeout time.Duration

	// How long would proxy try to dial server
	DialTimeout time.Duration
}

type ReverseProxy struct {

	// Metrics we track about this reverse proxy.
	metrics metrics.ProxyMetrics

	// Load balancer algorightm implementation
	loadBalancer loadbalance.Balancer

	// Rate limiter
	rateLimiter ratelimit.RateLimiter
}
