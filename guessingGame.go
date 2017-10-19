//Damian Gavin 19/10/17
//Web app in Go that prints "Guessing Game"

package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Guessing Game</h1>")
}

func main(){
	http.HandleFunc("/", server)//"/"handles any requests and passes to server
	http.ListenAndServe(":8080", nil)//webserver is started
}