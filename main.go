package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-host] [-port] [-interval] [-livenessport] [-insecure] [-nodename] COLLECTOR_ENDPOINT_URL\n", os.Args[0])
		flag.PrintDefaults()
	}

	var host string
	flag.StringVar(&host, "host", "localhost", "F@H Telnet Addr")
	var port string
	flag.StringVar(&port, "port", "36330", "F@H Telnet Port")
	var insecure bool
	flag.BoolVar(&insecure, "insecure", false, "Skip verifying collector's TLS cert")
	var interval time.Duration
	flag.DurationVar(&interval, "interval", 5*time.Second, "Send status once per $interval sec")
	var livenessPort string
	flag.StringVar(&livenessPort, "livenessport", "80", "Liveness Probe Port \"/healthz\"")
	var nodename string
	flag.StringVar(&nodename, "nodename", "", "Kubernetes Node Name or any other identifier (e.g., \"node123\")")

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	url := flag.Arg(0)
	timeout := 3 * interval

	// Liveness Probe
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	go func() {
		if err := http.ListenAndServe(":"+livenessPort, r); err != nil {
			log.Fatalf("Liveness Probe: http.ListenAndServe: %v", err)
		}
	}()

	// Telnet loop
	infoCh := StartTelnetInfoCh(host, port, interval)

	// POST loop
	t := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	c := &http.Client{
		Transport: t,
	}
	for {
		select {
		case <-time.After(timeout):
			log.Fatal("timed out")
		case s := <-infoCh:
			allRunning(s) // print log if any task is not running

			// inject data
			// - sc_nodename: Kubernetes node name (passed via arg)
			// - sc_hostname: Pod name (os.Hostname())
			hostname, err := os.Hostname()
			if err != nil {
				log.Printf("os.Hostname: %v", err)
				continue
			}
			s, err = injectValue(s, "sc_nodename", nodename)
			if err != nil {
				log.Printf("inject sc_nodename: %v", err)
				continue
			}
			s, err = injectValue(s, "sc_hostname", hostname)
			if err != nil {
				log.Printf("inject sc_hostname: %v", err)
				continue
			}

			// POST
			ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBufferString(s))
			if err != nil {
				log.Fatalf("http.NewRequestWithContext: %v", err)
			}
			h := http.Header{}
			h.Add("Content-Type", "application/json")
			req.Header = h
			resp, err := c.Do(req)
			if err != nil {
				log.Printf("http.Client.Do: %v", err)
				cancelFunc()
				continue
			}
			if resp.StatusCode != 200 {
				log.Printf("http.Client.Do: status=%d", resp.StatusCode)
				cancelFunc()
				continue
			}
			cancelFunc()
		}
	}
}
