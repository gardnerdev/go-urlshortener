# go-urlshortener


The goal of this exercise is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for **/bitcoin** to **youtube recording providing information about this cryptocurrency** we would look for any incoming web requests with the path **/bitcoin** and redirect them.


The next cryptocurrency is **/luna**, which will redirect us to official **Terra blockchain** youtube recording.

To complete this task I will need to implement the stubbed out methods in handler.go. 


## Usage
```
go run main.go
```