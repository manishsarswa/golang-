	package Test

	import (
		"database/sql"
		"reflect"
		"testing"
	)

	func Test(t *testing.T) {

		 testcases := []struct {
			input  Customer
			output Customer
		}{
			{Customer{id: 32, name: "Man", dob: "13dec",add: Address{id: 122,streetName: "wed",city:"wwww",cid: 32}},Customer{id: 32,name: "Man",dob: "13dec",add: Address{id: 122,streetName: "wed",city:"wwww",cid: 32}}},
			//{Customer{id: 6,name: "Manish",dob: "123dec",add: Address{id: 11,streetName: "wed",city:"wwww",cid: 6}},Customer{id: 6,name: "Manish",dob: "123dec",add: Address{id: 11,streetName: "wed",city:"wwww",cid: 6}}},
		}


		db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
		if err!=nil{
			panic(err)
		}

		defer db.Close()
		for i:=range testcases{
			var c Customer
			c=testCreateCustomer(db,testcases[i].input)
			if !reflect.DeepEqual(c,testcases[i].output){
				t.Errorf("FAIL")
			}
			t.Logf("PASS")
		}



	}
