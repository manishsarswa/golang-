package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)



type tst1 struct{
	req *http.Request
	output []byte
}
type customer struct {
	Name string
	Age int
	Address string
}
func TestNetwork1(t *testing.T) {
	//group:=tst{
	//	name:"Manish",
	//	age:20,
	//	address:"kormangala 5th cross",
	//}
	//
	//b,err:=json.Marshal(group)
	//c:= []byte(`{name:Manish, age:18, address:wert}`)
	//d := bytes.NewReader([]byte(`{"name":"Mans","age":20,"address":"werrty"}`))
	//fmt.Println(d)

	testcases:=[]struct{
		input []byte
		output []byte
	}{
		{[]byte(`{"name":"manish","age":20,"address":"kormangala"}`),[]byte(`{"name":"manish","age":20,"address":"kormangala"}`)},
		{[]byte(`{"name":"man","age":14,"address":"kogala"}`),[]byte(`not eligible`)},
		{[]byte(`{"name":"man","address":"koa"}`),[]byte(`not eligible`)},

	}

	for i:=range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodPost,"http://localhost:8080/customer/",bytes.NewBuffer(testcases[i].input))
		handler(w,req)
		resp := w.Result()

		body, _ := ioutil.ReadAll(resp.Body)
		if string(body)!=string(testcases[i].output){
			t.Error("Failed")
		}


	}
}
