package main

import "net/http"

func main() {


mux :=  http.NewServeMux() 
HandlerMux := mux

s := &http.Server{
	Addr:           ":8080",
	Handler:        HandlerMux,
}

s.ListenAndServe()

}