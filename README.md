# ShortyResty Coding Challenge - URL Shortener

## Overview
ShortyResty is a Golang REST API service that shortens URLs and allows the user to be redirected to the shortened URL.

## Purpose
The reason why URLs should be optimized is to rank better in search results and to be more convenient for users to type out. Also, the benefit of URL redirects is for the user to avoid 404 error pages and ensure the best experience when visiting a webpage

## Specifications

## Tests
Test ShortyResty.go using curl within the second terminal. <br>
<br>
Test Results of the GET request:

>    Trying 127.0.0.1... <br>
>  TCP_NODELAY set<br>
>  Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)<br>
>  GET /2SMtYvSg HTTP/1.1<br>
>  Host: 127.0.0.1:8080<br>
>  User-Agent: curl/7.64.1<br>
>  Accept: */*<br>
>  <br>
>  HTTP/1.1 200 OK<br>
>  Date: Fri, 18 Feb 2022 05:26:44 GMT<br>
>  Content-Length: 0<br>
>  <br>
>  Connection #0 to host 127.0.0.1 left intact<br>
>  Closing connection 0<br>


## Sources
[Gophercises](https://gophercises.com/)<br>
[LogRockets](https://blog.logrocket.com/creating-a-web-server-with-golang/)
