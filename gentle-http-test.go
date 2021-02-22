package main

import (
	"log"
	"net/http"
	//"os"
	//"runtime/pprof"
	"time"

	"github.com/levigross/grequests"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/transport"
)

var (
	httpsbase = "https://httpbin.org"
	httpbase  = "http://httpbin.org"
	path      = "/base64/bWFrZSBpdCBzbw=="
	iters     = 10
)

func main() {
	println("*** HTTPS")
	nethttp(iters, httpsbase)
	gentle(iters, httpsbase)
	gentletransport(iters, httpsbase)
	greqs(iters, httpsbase)
	println("*** HTTP")
	nethttp(iters, httpbase)
	gentle(iters, httpbase)
	gentletransport(iters, httpbase)
	greqs(iters, httpbase)
}

func greqs(iterations int, base string) {
	defer duration(track("grequests"))
	for i := 0; i < iterations; i++ {
		_, err := grequests.Get(base+path, nil)
		if err != nil {
			println(err.Error())
		}
	}
}

func nethttp(iterations int, base string) {
	defer duration(track("nethttp"))
	for i := 0; i < iterations; i++ {
		_, err := http.Get(base + path)
		if err != nil {
			println(err.Error())
		}
	}
}

func gentle(iterations int, base string) {
	cli := gentleman.New()
	cli.URL(base)
	defer duration(track("gentleman"))
	for i := 0; i < iterations; i++ {
		req := cli.Request().Path(path)
		_, err := req.Send()
		if err != nil {
			println(err.Error())
		}
	}
}

func gentletransport(iterations int, base string) {
	cli := gentleman.New()
	cli.URL(base)
	cli.Use(transport.Set(http.DefaultTransport))
	defer duration(track("gentleman transport"))
	for i := 0; i < iterations; i++ {
		req := cli.Request().Path(path)
		_, err := req.Send()
		if err != nil {
			println(err.Error())
		}
	}
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
