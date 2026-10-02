// Package app is the small amount of plumbing that api and worker share:
// connecting to postgres and serving the endpoints every pod must expose.
package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// Open connects to DATABASE_URL and waits until postgres actually answers.
// A pod can easily start before an operator-managed database is accepting
// connections, so the retry loop is what keeps the container from crash-looping.
func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	for i := 0; ; i++ {
		if err = db.Ping(); err == nil {
			return db, nil
		}
		if i == 30 {
			return nil, err
		}
		log.Printf("waiting for postgres: %v", err)
		time.Sleep(2 * time.Second)
	}
}

var (
	mu       sync.Mutex
	counters = map[string]int64{}
)

// Count bumps a counter exposed on /metrics.
func Count(name string, delta int64) {
	mu.Lock()
	counters[name] += delta
	mu.Unlock()
}

// Serve registers /health and /metrics on mux and blocks serving it.
//
// /health is what the kubelet probes. HEALTH_FAIL=true makes it answer 503 on
// purpose: that is the switch for shipping a deliberately broken release in
// Part 2 and watching the rollout refuse to finish.
func Serve(mux *http.ServeMux) error {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("HEALTH_FAIL") == "true" {
			http.Error(w, "unhealthy", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		names := make([]string, 0, len(counters))
		for name := range counters {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Fprintf(w, "# TYPE %s counter\n%s %d\n", name, name, counters[name])
		}
	})

	addr := ":" + Env("PORT", "8080")
	log.Printf("listening on %s", addr)
	return (&http.Server{Addr: addr, ReadHeaderTimeout: 5 * time.Second, Handler: mux}).ListenAndServe()
}

// Env reads an environment variable, falling back to def when it is unset.
func Env(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}
