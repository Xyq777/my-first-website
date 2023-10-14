package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/regReceive", regReceive)
	mux.HandleFunc("/reg", reg)
	mux.HandleFunc("/", login)
	mux.HandleFunc("/logReceive", logReceive)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
