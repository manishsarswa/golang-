package main

import (
	"database/sql"
	"reflect"
	"testing"
)


func Task(t *testing.T) {
	testcases:=[]struct{
		input int
		output []Customer
	}{
		{1,[]Customer{{id: 1,name: "manish",add: Address{id: 4,city: "raj",cid: 1}}}},
		{0,[]Customer{{id: 1,name: "manish",add: Address{id: 4,city: "raj",cid: 1}},{id: 2,name: "ishan",add: Address{id: 3,city: "up",cid: 2}},{id: 3,name: "bhava",add: Address{id: 2,city: "ap",cid: 3}},{id: 4,name: "pankaj",add: Address{id: 1,city: "up",cid: 4}}}},
	}

	db,err:=sql.Open("mysql", "root:Manish@123Sharma@/pratice")
	if err!=nil{
		panic(err)
	}
	for i:=range testcases{
		id:=testcases[i].input
		x:=handle(db,id)
		if !reflect.DeepEqual(x,testcases[i].output){
			t.Errorf("failed")
		} else{
			t.Logf("success")
		}

	}


}
