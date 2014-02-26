package service

import (
	"fmt"
	"github.com/wolfeidau/rifter/metrics"
)

type ServiceOptions struct {
	SslCertFile string
	SslKeyFile  string
	Port        string
	Endpoints   ListOptions
	Redis       string
	Auth        string
}

type ListOptions []string

func (o *ListOptions) String() string {
	return fmt.Sprint(*o)
}

func (lo *ListOptions) Set(value string) error {
	*lo = append(*lo, value)
	return nil
}

type Service struct {
	options *ServiceOptions
	//	proxy   *proxy.ReverseProxy
	metrics metrics.ProxyMetrics
}

func NewService(options *ServiceOptions) (*Service, error) {
	return &Service{options: options, metrics: metrics.NewProxyMetrics()}, nil
}

func (s *Service) Start() error {
	return nil
}
