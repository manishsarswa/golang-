package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type tst struct{
	req *http.Request
	out string
}

func TestServer(t *testing.T) {

	testcase:=[] struct{
		req *http.Request
		out string
	}{
		{httptest.NewRequest("GET", "http://localhost:8080/", nil), "Hello Golang"},
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

func TestServer1(t *testing.T){
	testcase:=[]tst{
		{httptest.NewRequest(http.MethodGet,"http://localhost:8080/print/",nil),"Printing Batman"},
	}

	str := []string{"abc", "def"}

	for i := range str {
		s := tst{httptest.NewRequest(http.MethodGet,fmt.Sprintf("http://localhost:8080/print/%s",str[i]),nil), fmt.Sprintf("Printing %s", str[i])}
		testcase = append(testcase, s)
	}

	for ind:=range testcase{
		w := httptest.NewRecorder()
		Printer(w, testcase[ind].req)
		resp := w.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		if string(body)!=testcase[ind].out{
			t.Error("Failed")
		}

	}
}
