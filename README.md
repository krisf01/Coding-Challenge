# ShortyResty Coding Challenge - URL Shortener

## Overview
ShortyResty is a Golang REST API service that shortens URLs and allows the user to be redirected to the shortened URL.

## Purpose
There are many advantages to shortened URLs as they are beneficial for businesses and are even great for aesthetic purposes. Who doesn't like a short and easy-to-type URL? Shortened URLs also rank better in search results and are more convenient for the user to type out. Along with a URL shortener, there is also a benefit to URL redirects. URL redirects are in place to allow the user to avoid 404 error pages and ensure the best experience when visiting a webpage.

## Specifications
* Create a web server on port 8080 using standard library "net/http" <br>
* The program then responds to a POST request, and the endpoint should be a POST method with URL/shorten <br>
  * The request should be in JSON as {“url”: “http://example.com/verylonguselessURLthatdoesnotseemtoend/123”}<br>
  * The response should be in JSON as {“short_url”: “http://127.0.0.1:8080/xxxxxxxx”} with an ID of a random string length of eight and be organized through the use of a map <br>
* The program then redirects the user when making a GET request, as the endpoint should be a GET method with URL/$ID. The request should be in the browser and or client to http://127.0.0.1:8080/ID and have 302 redirect the user to the long form of the URL

## Tests
Test ShortyResty.go using curl within the second terminal. <br>
<br>
Test Results of the GET request:
>   Trying 127.0.0.1... <br>
> TCP_NODELAY set <br>
> Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0) <br>
> GET / HTTP/1.1 <br>
> Host: 127.0.0.1:8080 <br>
> User-Agent: curl/7.64.1 <br>
> Accept: */* <br>
> <br>
> HTTP/1.1 200 OK <br>
> Date: Sat, 19 Feb 2022 01:15:43 GMT <br>
> Content-Length: 1 <br>
> Content-Type: text/plain; charset=utf-8 <br>
> <br>
> Connection #0 to host 127.0.0.1 left intact <br>
> Closing connection 0 <br>

<br>

Test Results of the POST request:
> TCP_NODELAY set<br>
> Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)<br>
> POST /shorten HTTP/1.1<br>
> Host: 127.0.0.1:8080<br>
> User-Agent: curl/7.64.1<br>
> Accept: */*<br>
> Content-Type: application/json<br>
> Content-Length: 72<br>
> <br>
> upload completely sent off: 72 out of 72 bytes<br>
> HTTP/1.1 201 Created<br>
> POST Request Content Type: : application/json<br>
> Date: Sat, 19 Feb 2022 01:19:01 GMT<br>
> Content-Length: 47<br>
> Content-Type: text/plain; charset=utf-8<br>
> <br>
> http://127.0.0.1:8080/bPlNFGdS<br>
> Connection #0 to host 127.0.0.1 left intact<br>
> Closing connection 0<br>

<br>

Test Results of the Unit Test for HTTP Handler:
> Running tool: /usr/local/go/bin/go test -timeout 30s -run ^TestTeapotHandler$

## Resources
* [Gophercises: URL Shortener](https://gophercises.com/)<br>
* [LogRockets: Creating a Web Server with Golang](https://blog.logrocket.com/creating-a-web-server-with-golang/)<br>
* [Intersog: How to Make a Custom URL Shortener Using Golang and Redis](https://intersog.com/blog/how-to-write-a-custom-url-shortener-using-golang-and-redis/)<br>
* [bindersfullofcode: Writing a URL Shortener in Go](http://bindersfullofcode.com/2019/02/12/golang-url-shortener.html)<br>
* [The Polyglot Developer: Create a URL Shortener with Golang and Couchbase NoSQL](https://www.youtube.com/watch?v=OVBvOuxbpHA)<br>
* [Documenting APIs: curl into and installation](https://idratherbewriting.com/learnapidoc/docapis_install_curl.html)
* [Learn Go with Tests: Revisiting HTTP Handlers](https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited)
* [LogRockets: A Deep Dive into Unit Testing in Go](https://blog.logrocket.com/a-deep-dive-into-unit-testing-in-go/)
