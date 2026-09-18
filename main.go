package main

import (
	"net/http"
	"sync/atomic"
	"fmt"
)


type apiConfig struct {
		fileserverHits atomic.Int32
	}

	//definesc functii pe acel struct.
	func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cfg.fileserverHits.Add(1)
			next.ServeHTTP(w, r)
		})
	}

	//definesc functii pe acel struct.
	func (cfg *apiConfig) countHits() string {
		return fmt.Sprintf("Hits: %d",cfg.fileserverHits.Load() )
	}

	//definesc functii pe acel struct.
	func  (cfg *apiConfig) ResetHits() {
		 	cfg.fileserverHits.Store(0) 
	}

func main() {

	cfg := &apiConfig{}

	mux := http.NewServeMux()

		// Readiness endpoint
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Metrics endpoint
	mux.HandleFunc("GET /metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cfg.countHits()))
	})

	mux.HandleFunc("POST /reset",func(w http.ResponseWriter, r *http.Request)  {

			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			cfg.ResetHits()
			w.Write([]byte("OK"))
			
		})

	// FileServer
	fs := http.FileServer(http.Dir("."))

	// FileServer-ul este acum disponibil la /app/
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app", fs)))


	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}