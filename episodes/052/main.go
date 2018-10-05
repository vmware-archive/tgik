package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/kubicorn/kubicorn/pkg/logger"
	"time"
	"math/rand"
)

var metricsAddr = flag.String("metrics-address", ":1313", "The address to listen on for metrics requests.")
var httpAddr = flag.String("http-address", ":80", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	logger.Always("Listening [metrics] on address %s", *metricsAddr)
	http.Handle("/metrics", promhttp.Handler())
	logger.Always("register [/metrics]")
	logger.Always("Serving metrics...")


	// This will literally bring your cluster down
	//go func() {
	//	for {
	//		go func() {
	//			time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	//			}()
	//	}
	//}()


	log.Fatal(http.ListenAndServe(*metricsAddr, nil))

	//logger.Always("Listening [http] on address %s", *httpAddr)
	//router := mux.NewRouter()
	//logger.Always("register [/]")
	//router.HandleFunc("/", homeHandler)
	//logger.Always("Serving HTTP...")
	//
	//
	//srv := &http.Server{
	//	Handler:      router,
	//	Addr:         *httpAddr,
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//log.Fatal(srv.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to TGIK"))
	w.WriteHeader(200)
}