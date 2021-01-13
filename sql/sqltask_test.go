package TestSql

import (
	"database/sql"

	"reflect"
	"testing"
)
//type customers struct {
//	ID int
//	Name string
//	DOB string
//	add address
//}
//type address struct{
//	ID int
//	street_Name string
//	city string
//	state string
//}

func TestSql(t *testing.T) {

	testcases:=[]struct{
			input int
			output []Customer
	}{
		{1,[]Customer{{ID:1,name:"manish",DOB:"12dec",add:address{ID: 3,street_Name: "5th cross",city: "bangalore",state: "karnataka",customer_ID:1}}}},
		{2,[]Customer{{ID:2,name:"deepika",DOB:"14may",add:address{ID:4,street_Name: "5th cross 5th block",city: "bangalore",state: "karnataka",customer_ID: 2}}}},
		{0,[]Customer{{ID:3,name:"ishan",DOB:"23dec",add:address{ID:1,street_Name: "zopsmart",city: "bangalore",state: "karnataka",customer_ID: 3}},{ID:4,name:"sharma",DOB:"2jan",add:address{ID:2,street_Name: "zopsmart 24th main",city: "bangalore",state: "karnataka",customer_ID: 4}},{ID:1,name:"manish",DOB:"12dec",add:address{ID: 3,street_Name: "5th cross",city: "bangalore",state: "karnataka",customer_ID:1}},{ID:2,name:"deepika",DOB:"14may",add:address{ID:4,street_Name: "5th cross 5th block",city: "bangalore",state: "karnataka",customer_ID: 2}}}},
	}
	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
	if err!=nil{
		panic(err)
	}

	//defer db.Close()
	for i:=range testcases{
		a:=testcases[i].input
		var x[]Customer
		x=getcustomer(db,a)

		if !reflect.DeepEqual(x,testcases[i].output){
			t.Fatalf("%v failed",i+1)
		}
		t.Logf("success")

	}
}
