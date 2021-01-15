package TestSql

import (
	"database/sql"

	"reflect"
	"testing"
)

func TestSql(t *testing.T) {

	testcases:=[]struct{
			input string
			output []Customer
	}{
		{"1",[]Customer{{id:1,name:"manish",DOB:"12dec",add:Address{id: 3,streetName: "5th cross",city: "bangalore",state: "karnataka",customerId:1}}}},
		{"2",[]Customer{{id:2,name:"deepika",DOB:"14may",add:Address{id:4,streetName: "5th cross 5th block",city: "bangalore",state: "karnataka",customerId: 2}}}},
		{"0",[]Customer{{id:3,name:"ishan",DOB:"23dec",add:Address{id:1,streetName: "zopsmart",city: "bangalore",state: "karnataka",customerId: 3}},{id:4,name:"sharma",DOB:"2jan",add:Address{id:2,streetName: "zopsmart 24th main",city: "bangalore",state: "karnataka",customerId: 4}},{id:1,name:"manish",DOB:"12dec",add:Address{id: 3,streetName: "5th cross",city: "bangalore",state: "karnataka",customerId:1}},{id:2,name:"deepika",DOB:"14may",add:Address{id:4,streetName: "5th cross 5th block",city: "bangalore",state: "karnataka",customerId: 2}}}},
		{"4555",[]Customer(nil)},
		{"1 or 1=1", []Customer(nil)},
	}
	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
	if err!=nil{
		panic(err)
	}

	defer db.Close()
	for i:=range testcases{
		id:=testcases[i].input

		rows:=getCustomer(db,id)

		if !reflect.DeepEqual(rows,testcases[i].output){
			t.Fatalf("FAIL")
		}
		t.Logf("PASS")

	}
}
