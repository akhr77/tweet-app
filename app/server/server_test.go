package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRouter(t *testing.T) {
	expected := ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer ts.Close()

	res, err := http.DefaultClient.Get(ts.URL + "/post")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(res)
	defer res.Body.Close()

	if !reflect.DeepEqual(res.StatusCode, 200) {
		t.Errorf("http.Get returnCode actual:%v, expected:%v", res.StatusCode, 200)
	}
}
