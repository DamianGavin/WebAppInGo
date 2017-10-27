# Problem sheet 2
Data Representation and Querying

This is our 2nd problem sheet entitled "Web applications" There are 8 parts
to it and I will go through them individually.

### Some of the resources that I used

1. http://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
2. https://curl.haxx.se/docs/httpscripting.html
3. https://stackoverflow.com/questions/28793619/golang-what-to-use-http-servefile-or-http-fileserver
4. https://data-representation.github.io/
5. https://getbootstrap.com/docs/4.0/getting-started/introduction/#starter-template
6. https://www.youtube.com/watch?v=GTSq1VPPFco&feature=youtu.be
7. https://golang.org/pkg/text/#Template.Execute
8. https://astaxie.gitbooks.io/build-web-application-with-golang/de/04.1.html

### Running the code

First you must have "Go" on your computer. This can be found and downloaded at https://golang.org/.
If you navigate to the folder that contains these files you can type :
go run <file name> to run the code.

## Part One: Guessing game

ToDo::Create a web application in Go that responds with the text “Guessing game”. This should be the response body irrespective of what request is received. Explain in your README how to examine the response, including the headers, using curl.

What’s going on here is that in the main function, we are handling any requests made to "/" by passing these requests on to a function called server. The http package that has been imported contains the function http.HandleFunc() which is taking care of sending all the requests made to "/" on to our helloWorld function. The http.HandleFunc() takes two inputs, the first being a pattern which is a string and the second is a handler (a function that needs a ResponseWriter and a pointer to a Request). In order to satisfy the http.HandleFunc our helloWorld function receives two inputs, the first input is called w and is of type http.ResponseWriter. W is where a response can be sent. The second input to the helloWorld function is called r and this is a pointer to a http.Request (the * prefix indicates that this is a pointer

When I opened cmder, where I already have curl installed in my bin folder I could examine the response, including the headers by this command

```
curl --verbose localhost:8080
```

This was the result

```
 Rebuilt URL to: localhost:8080/                          
   Trying ::1...                                          
 TCP_NODELAY set                                          
 Connected to localhost (::1) port 8080 (#0)              
 GET / HTTP/1.1                                           
 Host: localhost:8080                                     
 User-Agent: curl/7.55.1                                  
 Accept: */*                                              
                                                          
 HTTP/1.1 200 OK                                          
 Date: Thu, 19 Oct 2017 12:54:59 GMT                      
 Content-Length: 13                                       
 Content-Type: text/plain; charset=utf-8                  
                                                          
uessing Game* Connection #0 to host localhost left intact 
```
If I want only the server response's HTTP headers, instead of the page data I can use 

```
curl -I localhost:8080
```




## Part Two-Make the text a H1

For this problem I simply inserted a h1 tag in the print statement

```
fmt.Fprintf(w, "<h1>Guessing Game</h1>")
```

## Part Three-Serve a page using Bootstrap

Change the web application to serve a web page rather than hard-coding the text into the web application executable. Use the Bootstrap starter template,changing the header to say "Guessing game". Add a link on the page to the relative URL /guess with the text “New game”. Have this page served as the root resource in the web application.

For this part I opened the provided bootstrap template link and copied the html into a new file in the project folder. I basically needed to swap my "Guessing Game" print line with the html template. In my .go file I inserted this line instead
```
func server(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r, "guess.html")}
```
The ServeFile is a lower level helper that can be used to implement something similar to FileServer, or implement my own path munging potentially, and any number of things. It simply takes the named local file and sends it over the HTTP connection.

```
 <h1>Guessing Game</h1><a href = "/guess">New game</a>
```
This is the link I inserted in my html file, It is a basic html tag and renders the words "New game" as a link to /guess which appears in the url. 
And repeat


## Part Four-Add a guess route

Create a new route in your application at /guess. Have it serve a new page called guess.html. Use the same Bootstrap code for the page as in index.html but add a level 2 heading with the text “Guess a number between 1 and 20”. Add a form, with a text input with the name “guess” and a submit button with the label “Guess”. The action of the form should be /guess and the method should be GET.

Here I inserted this code into my guess.html

```
<h2>Guess a number between 1 and 20</h2>
	
	<form action = "/guess" method = "GET">
	Enter guess:
	
	<input name = "guess" type = "text">
	<br/>
	<br/>
	<button value = "guess">Guess</button>
	</form>
```
### Part Five-Turn the guess page into a template

Change the web application to use the guess.tmpl file as a template. Add a single variable to the template called Message and create a struct in Go containing a single string. Create an instance of the struct with the string set to “Guess a number between 1 and 20” and have the template render this in the H2 tag.

Here I copied the bootstrap code, but rather than giving it the html extension I called it guess.tmpl. This file will only render internally and cannot be accessed by the user. It has all the functionality of the other html files but the internal mechanisms of the program will use it only as a template.
I created the following struct in the .go file

```
type templateData struct{
Message string
}
```
and this was my variable declaration, also in .go.
```
message :="Please guess a number between 1 and 20"
```
I then had to insert
```
t.Execute(w, &templateData{Message:message})
```
within my func guesshandler. Then I needed to call it using data binding in my guess.tmpl within the required H2 tag.
```
 <h2>{{.Message}}</h2>
```

### Part Six-Check for a cookie

In the /guess handler check if the cookie called target is set. If it is not, generate a random number between 1 and 20 and set a cookie called target to that value in the response. Otherwise, leave the cookie at its current value.

```
var cookie, err = r.Cookie("target")
		
			if err == nil{
				
				target, _ = strconv.Atoi(cookie.Value)
				if target ==0{
					target = rand.Intn(1-20)
				}
			}
		
			cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)
```

## Part Seven-Check for a variable

Have the /guess handler check if a URL encoded variable called guess has been set to an integer. If it has, have the text “You guessed {guess}.” inserted into the template where {guess} is replaced with the value of guess.

Here I needed to change the input type from "text" to "number" in my .tmpl template. I also instigated my data binding metod to receive the guessed number as {{.Guess}}, I just used a p tag.

```
<input type="number" name="guess">
```
```
<p>You guessed {{.Guess}}</p>
```
```
type templateData struct{
Message string
Guess string
}
```
I had to change my struct to accomodate guess, and also read guess as an int
```
r.ParseForm();
guess := r.FormValue("guess")
```

## Part Eight-Compare the cookie and the guess

If the target cookie and the guess variable are both set, then have the /guess handler compare them. If they are equal, set the target cookie to another randomly generated integer, and have the template display a congratulations message and a link to create a new game. Otherwise, have the template display a message telling the user what their guess was and whether it was too high or too low.

```
Give an example
```

## Part Nine-Use the POST method

Change your HTML form in guess.html to use the POST method instead. Make sure your application still works, bug fixing it if necessary.

```
Give an example
```

## Built With

## Contributing

