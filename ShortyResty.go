package main
 
import (
   "log"
   "math/rand" // for rand
   "net/http" // JSON for request/response
   "encoding/json"
   "strings"
   "net/url"
)
 
// JSON response for shortened URLs
type jsonResponse struct {
   ShortURL string `json:"short_url"`
}
 
// JSON redirect for long URLs
type jsonRedirect struct {
   Long_url string `json:"url"`
}
 
var pathToUrls map[string]jsonRedirect
const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
 
// function for the shorten function to shorten URL and contain an ID with random string of eight
func string_eight() string {
   // random characters for shortened URLs
   var id string = ""
   for i := 0; i < 8; i++ {
       id += string(chars[rand.Intn(len(chars))])
   }
   // returning random id char
   return id
}
 
// Redirect function handles the GET requests and will redirect the user
func redirect(writer http.ResponseWriter, req *http.Request) {

   // leaving id of the URL only
   var id string = strings.TrimPrefix(req.RequestURI, "/") 
   if _, ok := pathToUrls[id]; ok {
	   http.Redirect(writer, req, pathToUrls[id].Long_url, http.StatusFound)
   }
}
 
// Shorten long URL function. Uses a map for redirect endpoint
func shorten(writer http.ResponseWriter, req *http.Request){

	// header
	writer.Header().Set( "","application/json")
	writer.WriteHeader(http.StatusCreated)

	// decoding the jsn redirect for long URLs
	var longUrl jsonRedirect 
	json.NewDecoder(req.Body).Decode(&longUrl)

	// checking if URL is valid
	_, err := url.ParseRequestURI(longUrl.Long_url)
	if err != nil {
		http.Error(writer, "Error, not valid!", http.StatusBadRequest)  
		return
	}

	// variable declaration for string and if it exists comparison
	var id string
	var exists bool = true

	// While id exists in the map, it will generate random IDs
	for exists {
		id = string_eight()
		if _, key := pathToUrls[id]; !key {
			exists = false
		}
	}
	// Adding the id of the long URL to the map
	pathToUrls[id] = longUrl  

	// creating the shortened url
	var newUrl string
   	newUrl = "http://127.0.0.1:8080/" + (string(id))
	shortUrl := jsonResponse{newUrl}

	// creating JSON response so it can be sent
	json.NewEncoder(writer).Encode(shortUrl) 
}

// starting web server within the main function
func main() {
   // creating the URL map
   pathToUrls = make (map[string]jsonRedirect)

   // handle functions
   http.HandleFunc("/", redirect)
   http.HandleFunc("/shortened", shorten)

   // creating web server on port 8080
   if err := http.ListenAndServe(":8080", nil); err != nil {
       log.Fatal(err) // checking for valid website
}
}

