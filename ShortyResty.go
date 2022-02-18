package main
 
import (
   "log"
   "math/rand" // for rand
   "net/http" // JSON for request/response
   "encoding/json"
   "strings"
)
 
// JSON response for shortened URLs
type jsonResponse struct {
   ShortURL string `json:"short_url"`
}
 
// long URLs
type jsonRedirect struct {
   Long_url string `json:"url"`
}
 
var pathToUrls map[string]jsonRedirect
 
// function for shortened URL to contain an ID with random string of eight
func string_eight() string {
 
   // random characters for shortened URls
   var stringRand string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
   var id string = ""
   for i := 0; i < 8; i++ {
       id += string(stringRand[rand.Intn(len(stringRand))])
   }
 
   //adding random characters to end of url
   var newUrl string
   newUrl = "http://127.0.0.1:8080/" + (string(id))
   return newUrl
}
 
// function for redirecting to the long for of the URL
func handleRedirect(writer http.ResponseWriter, req *http.Request) {
   var id string = strings.TrimPrefix(req.RequestURI, "/")       // Trims prefix from request URI, leaving just the ID
   if _, key := pathToUrls[id]; key {
       http.Redirect(writer, req, pathToUrls[id].Long_url, 302)       // 302 redirect to the long URL
   } else {
       writer.Write([]byte(req.RequestURI + " is not linked to a long URL."))       // Respond with error message if id is not in urlMap
       return
   }
}
 
// starting web server
func main() {
 
   // creating the url map
   pathToUrls = make (map[string]jsonRedirect)
 
   mux := http.NewServeMux()
   mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
       if req.Method == http. MethodGet {
           // the url with the random id attacthed
           newUrl := string_eight()
           // random url
           data := jsonResponse {
               ShortURL: newUrl,
           }
           json.NewEncoder(writer).Encode(data)
           // the redirecting function
           http.HandleFunc("/", handleRedirect)
       }
   })
 
   // creating web server on port 8080
   if err := http.ListenAndServe(":8080", nil); err != nil {
       log.Fatal(err) // checking for valid website
}
}

