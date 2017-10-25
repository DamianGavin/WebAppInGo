//Damian Gavin 19/10/17
//Web app in Go that prints "Guessing Game"

package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
	"strconv"
)

type myMsg struct{
Message string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","text/html")

	http.ServeFile(w, r, "04_index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request){

	//http.ServeFile(w, r, "04_guess.html")

		message :="Guess a number between 1 and 20"
		
			rand.Seed(time.Now().UTC().UnixNano())
		
			target:=0
			var cookie, err = r.Cookie("target")
		
			if err == nil{
				
				target, _ = strconv.Atoi(cookie.Value)
				if target ==0{
					target = rand.Intn(20-1)
				}
			}
		
			cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)
			
			t, _ := template.ParseFiles("04_guess.tmpl")

			t.Execute(w, &myMsg{Message:message})
}


func main(){
	http.HandleFunc("/", server)//"/"handles any requests and passes to server
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)//webserver is started
}