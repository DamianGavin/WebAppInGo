package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Guessing Game")
}

func main(){
	http.HandleFunc("/", server)
	http.ListenAndServe(":8080", nil)
}