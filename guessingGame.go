//Damian Gavin 19/10/17
//Web app in Go that prints "Guessing Game"

package main

import (
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r, "guess.html")}

func main(){
	http.HandleFunc("/", server)//"/"handles any requests and passes to server
	http.ListenAndServe(":8080", nil)//webserver is started
}