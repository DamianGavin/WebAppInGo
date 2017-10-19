# Problem sheet 2
Data Representation and Querying

This is our 2nd problem sheet entitled "Web applications" There are 8 parts
to it and I will go through them individually.

### Some of the resources that I used

http://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
https://curl.haxx.se/docs/httpscripting.html

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

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Part Four

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
<<<<<<< HEAD
=======

>>>>>>> 49edb1f94a3a386aeb3d648fbd90b0223cf12880
