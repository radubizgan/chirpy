package main

import "net/http"

func main() {



mux :=  http.NewServeMux() 
fs := http.FileServer(http.Dir("."))// servește fișiere din directorul curent
pictures := http.FileServer(http.Dir("assets"))// servește fișiere din directorul indicat
mux.Handle("/",fs )  // toate requesturile care încep cu "/" sunt gestionate de fs
mux.Handle("/assets",pictures )  // toate requesturile care încep cu "/" sunt gestionate de fs


s := &http.Server{
	Addr:           ":8080",
	Handler:        mux,
}

s.ListenAndServe()

}