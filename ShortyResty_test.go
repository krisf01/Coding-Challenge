package ShortyResty

import (
   "testing" // unit testing
   "net/http/httptest" // unit testing
    "net/http"
)

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