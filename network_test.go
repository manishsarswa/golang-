package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	testcase:=[] struct{
		req *http.Request
		out string
	}{
		{httptest.NewRequest("GET", "http://192.168.1.55:8080/", nil), "Hello Golang"},
	}

	for ind:=range testcase{
		w := httptest.NewRecorder()
		handle(w, testcase[ind].req)
		resp := w.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		if string(body)!=testcase[ind].out{
			t.Error("Failed")
		}

	}


}
