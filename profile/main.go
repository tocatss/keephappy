package main

// TODO: add readme.

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"strings"
	"time"
)

var buff []string

func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", http.HandlerFunc(welcomeHandler))

	registerPprof(mux)
	memoryLeak()
	http.ListenAndServe(":9090", mux)
}

func memoryLeak() {
	go func() {
		for {
			buff = append(buff, strings.Repeat("spp", 100000))
			time.Sleep(1 * time.Second)
		}
	}()
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func registerPprof(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
