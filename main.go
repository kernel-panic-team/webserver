package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	sch := gocron.NewScheduler(time.UTC)
	_, err := sch.Every(1).Day().At("00:00").Do(worker)
	if err != nil {
		log.Fatalf("unable to schedule task: %v", err)
	}
	sch.StartAsync()
	defer sch.Stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", getCheckHealthFunc())
	go func() { _ = http.ListenAndServe(":8080", mux) }()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-ch

	log.Printf("service stopped")
}

func worker() {
	// do something
}

func getCheckHealthFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
