package handler_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	// "bitbucket.org/skshahriarahmed/sh_ra/route"
// 	"github.com/gorilla/mux"
// )


// func TestLogin(t *testing.T) {
// 	// go main()
// 	r := mux.NewRouter()
// 	ts := httptest.NewServer(r)
// 	defer ts.Close()

// 	res, err := http.Get(ts.URL + "/login")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if res.StatusCode != http.StatusOK {
// 		t.Fatalf("Status code is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
// 	}
// }



// func NewApplicationHandler() http.Handler {
//     mux := mux.NewRouter()
//     mux.HandleFunc("/", handler)

//     return mux
// }