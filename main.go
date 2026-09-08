package main

import "net/http"

func main() {



mux :=  http.NewServeMux() 
fs := http.FileServer(http.Dir("."))// servește fișiere din directorul curent
mux.Handle("/",fs )  // toate requesturile care încep cu "/" sunt gestionate de fs

s := &http.Server{
	Addr:           ":8080",
	Handler:        mux,
}

s.ListenAndServe()

}