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
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {

		// 1. Setăm Content-Type
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		// 2. Trimitem status code 200
		w.WriteHeader(http.StatusOK)

		// 3. Trimitem body-ul
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/metrics",func(w http.ResponseWriter, r *http.Request)  {

		// 1. Setăm Content-Type
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		// 2. Trimitem status code 200
		w.WriteHeader(http.StatusOK)

		// 3. Trimitem body-ul
		w.Write([]byte(cfg.countHits() ))
		
	})

	mux.HandleFunc("/reset",func(w http.ResponseWriter, r *http.Request)  {

			// 1. Setăm Content-Type
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")

			// 2. Trimitem status code 200
			w.WriteHeader(http.StatusOK)

			cfg.ResetHits()

			// 3. Trimitem body-ul
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