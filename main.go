package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/wolfeidau/rifter/service"
	"os"
	"runtime"
)

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	options := &service.ServiceOptions{}

	flag.StringVar(&options.Port, "port", ":1883", "The IP Address and TCP port to listen")
	flag.StringVar(&options.SslCertFile, "sslcert", "", "File containing SSL Certificates")
	flag.StringVar(&options.SslKeyFile, "sslkey", "", "File containing SSL Private Key")
	flag.Var(&options.Endpoints, "mqtt", "MQTT endpoints")
	flag.Parse()

	svc, err := service.NewService(options)
	if err != nil {
		glog.Fatalf("Failed to init service, error:", err)
	}

	glog.Fatal(svc.Start())
}
