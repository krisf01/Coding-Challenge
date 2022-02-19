package main

import (
   "log"
   "math/rand" // for rand
   "net/http" // JSON for request/response
   "encoding/json"
   "strings"
   "net/url"
   "testing" // unit testing
   "net/http/httptest" // unit testing
)
 

// JSON response for shortened URLs
type jsonResponse struct {
   ShortURL string `json:"short_url"`
}
 
// JSON redirect for long URLs
type jsonRedirect struct {
   LongURL string `json:"url"`
}

// Map
var pathToUrls map[string]jsonRedirect
 
// Function to shorten the U              RL and contain an ID with random length of eight
func string_eight() string {
	// Characters for the ID
	var variables = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

	// Randomizing the characters into length of eight
   var id string = ""
   for i := 0; i < 8; i++ {
       id += string(variables[rand.Intn(len(variables))])
   }

   // Returning the random variable ID
   return id
}
 
// Redirect function handles the GET requests and will redirect the user
func redirect(writer http.ResponseWriter, req *http.Request) {

   // Leaving id of the URL only
   var id string = strings.TrimPrefix(req.RequestURI, "/") 

   // Checking whether or not ID is within the map. Redirects the user, otherwise error
   if _, key := pathToUrls[id]; key {
		http.Redirect(writer, req, pathToUrls[id].LongURL, http.StatusFound)
		return
	} else {
		writer.Write([]byte(req.RequestURI + ""))
        return
	}
}
 
// Shorten long URL function. Uses a map for redirect endpoint.
func shorten(writer http.ResponseWriter, req *http.Request){

	// Label HTTP header and request the header to be returned
	writer.Header().Set("POST Request Content Type: ", "application/json")
	writer.WriteHeader(http.StatusCreated)

	// Decode the JSON redirect for long URLs. `json:"url"`
	var long_Url jsonRedirect 
	json.NewDecoder(req.Body).Decode(&long_Url)

	// Checking if the provided URL is valid
	_, err := url.ParseRequestURI(long_Url.LongURL)
	if err != nil {
		http.Error(writer, "Error, URL is not valid!", http.StatusBadRequest)  
		return
	}

	// While id exists in the map, it will generate random IDs
	var exists bool = true
	var id string

	for exists {
		id = string_eight()
		if _, key := pathToUrls[id]; !key {
			exists = false
		}
	}

	// Adding the id of the long URL to the map
	pathToUrls[id] = long_Url  

	// Creating the shortened url
	var new_Url string
   	new_Url = "http://127.0.0.1:8080/" + id
	short_Url := jsonResponse{new_Url}

	// Creating JSON response so it can be sent
	json.NewEncoder(writer).Encode(short_Url) 
}

// Starting web server within the main function
func main() {
   // Creating the URL map
   pathToUrls = make (map[string]jsonRedirect)

   // Handle functions for redirect and shorten
   http.HandleFunc("/", redirect)
   http.HandleFunc("/shorten", shorten)

   // Creating web server on port 8080
   if err := http.ListenAndServe(":8080", nil); err != nil {
       log.Fatal(err) // checking for valid website
	}
}

// Unit Test Function for HTTP Handler
func Teapot(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusTeapot)
}

// Unit Test Function for HTTP Handler
func TestTeapotHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}
