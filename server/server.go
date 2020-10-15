package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", WelcomeHandler)
	server := &http.Server{
		Addr:    ":0112",
		Handler: mux,
	}

	done := make(chan interface{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		osCall := <-c
		log.Printf("system call: %+v", osCall)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("shutdown err %v", err)
		}
		close(done)
	}()

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
		log.Print("servier is closed")
	}

	<-done
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Status", strconv.Itoa(http.StatusOK))
	fmt.Fprint(w, "ok")
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Middleware", "Middleware")
		next.ServeHTTP(w, r)
	})
}
