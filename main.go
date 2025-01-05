package main

import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Auto sync enabling v0.0.18</h1>"))
	})
	http.ListenAndServe(":8181", nil)
}