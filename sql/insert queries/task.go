package Test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct{
	id int
	name string
	dob string
	add Address
}

type Address struct{
	id int
	streetName string
	city string
	cid int
}
func testCreateCustomer(db *sql.DB,c Customer) Customer {
	var idr[] interface {}
	var addr[] interface{}

	query:=`insert into c1 value(?,?,?)`

	idr=append(idr,c.id)
	idr=append(idr,c.name)
	idr=append(idr,c.dob)
	_,err:=db.Query(query,idr...)
	if err!=nil{
		fmt.Println("err in first", err)
	}
	query=`insert into c2 value(?,?,?,?)`
	addr=append(addr,c.add.id)
	addr=append(addr,c.add.streetName)
	addr=append(addr,c.add.city)
	addr=append(addr,c.add.cid)

	_,er:=db.Query(query,addr...)
	if er!=nil{
		fmt.Println("err in second",er)
	}
	query=`select * from c1 inner join c2 on c1.id=c2.cid where c1.id=?`
	var idd [] interface{}
	idd=append(idd,c.id)
	rows,err:=db.Query(query,idd...)
	if err != nil {
		fmt.Println("err in third",err)
	}

	var cust Customer
	for rows.Next() {
		err = rows.Scan(&cust.id, &cust.name, &cust.dob, &cust.add.id, &cust.add.streetName, &cust.add.city, &cust.add.cid)
		if err != nil {

		}


	}

	return cust




}
