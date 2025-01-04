package main

import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Hello world Lucas</h1>"))
	})
	http.ListenAndServe(":8181", nil)
}