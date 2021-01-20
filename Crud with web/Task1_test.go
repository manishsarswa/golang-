package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//type Define struct{
//	Input string
//	Output []Customer
//}
//func Test(t *testing.T) {
//
//	testcases:=[]struct{
//		input string
//		output []Customer
//	}{
//		{"?name=manish", []Customer{{4, "manish", "29mar", "kormangala"}}},
//		{"?name=manish",[]Customer(nil)},
//	}
//
//	for i:= range testcases{
//		w:=httptest.NewRecorder()
//		req:=httptest.NewRequest(http.MethodGet,"/customer"+testcases[i].input,nil)
//		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].input})
//		handleGet(w,req)
//
//		var c[] Customer
//		err:=json.Unmarshal(w.Body.Bytes(),&c)
//		if err!=nil{
//			t.Log(err)
//		}
//		fmt.Println("test got ",c)
//
//		if !reflect.DeepEqual(c,testcases[i].output){
//			t.Errorf("Failed in %v testcase",i)
//		}
//	}
//}
//
//
//func TestGetId(t *testing.T){
//	testcases:=[]Define{
//		{"1",[]Customer{{1,"manish","12dec","Bikaner"}}},
//		{"0",[]Customer{{1,"manish","12dec","Bikaner"}}},
//		{"1234",[]Customer(nil)},
//	}
//
//	for i:=range testcases{
//		req:=httptest.NewRequest(http.MethodGet,"/customer/"+testcases[i].Input,nil)
//		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].Input})
//		w:=httptest.NewRecorder()
//		handleGetid(w,req)
//
//
//
//		var result []Customer
//		err:=json.Unmarshal(w.Body.Bytes(),&result)
//		if err!=nil{
//			t.Log(err)
//		}
//
//
//		if !reflect.DeepEqual(result,testcases[i].Output){
//			t.Errorf("Expected result  is %v and got result is %v",result,testcases[i].Output)
//		}
//
//	}
//
//
//}
//
//
//func TestPostid(t *testing.T) {
//	testcases:=[]struct{
//		input []byte
//		output []Customer
//	}{
//		{[]byte(`{"Id":10,"Name":"sharma","Dob":"13dec","Address":"Rajasthan"}`),[]Customer{{10,"sharma","13dec","Rajasthan"}}},
//		{[]byte(`{"Id":11,"Name":"Ankit","Dob":"23jan","Address":"delhi"}`),[]Customer{{11,"Ankit","23jan","delhi"}}},
//		{[]byte(`{"Id":12,"Name":"aniket","Dob":"14feb","Address":"bangalore"}`),[]Customer{{12,"aniket","14feb","bangalore"}}},
//
//	}
//
//
//	for i:=range testcases{
//		req:=httptest.NewRequest(http.MethodPost,"/customer/",bytes.NewBuffer(testcases[i].input))
//		w:=httptest.NewRecorder()
//		handlePost(w,req)
//		var c []Customer
//		json.Unmarshal(w.Body.Bytes(),&c)
//		if !reflect.DeepEqual(c,testcases[i].output){
//			t.Errorf("Failed in %v testcase",i)
//		}
//
//	}
//}
//

func TestPut(t *testing.T){
	testcases:=[]struct {
		input string
		body []byte
		output []Customer
	}{
		{"4",[]byte(`{"Name":"manishsharma"}`),[]Customer{{4,"manishsharma","29mar","kormangala"}}},

	}


	for i :=range testcases{

		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodPut,"/customer?id="+string(testcases[i].input),bytes.NewBuffer(testcases[i].body))
		updateCustomer(w,req)
		var c Customer
		json.Unmarshal(w.Body.Bytes(),&c)
		if !reflect.DeepEqual(c,testcases[i].output){
			t.Errorf("Failed in %v testcase ",i)
		}
	}
}
//func TestDelete(t *testing.T) {
//	testcases:=[]Define{
//		{"10",[]Customer{{10,"sharma","13dec","Rajasthan"}}},
//		{"0",[]Customer(nil)},
//		{"1234",[]Customer(nil)},
//	}
//
//
//	for i :=range testcases{
//		w:=httptest.NewRecorder()
//		req:=httptest.NewRequest(http.MethodDelete,"/customer/"+testcases[i].Input,nil)
//		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].Input})
//		handleDelete(w,req)
//		var c []Customer
//		json.Unmarshal(w.Body.Bytes(),&c)
//		//fmt.Println(c)
//		//fmt.Println(testcases[i].Output)
//		if !reflect.DeepEqual(c,testcases[i].Output){
//			t.Errorf("Failed in %v testcase ",i)
//		}
//
//	}
//}


