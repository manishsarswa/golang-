package TestSql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	id  int
	name string
	DOB  string
	add  Address
}
type Address struct {
	id         int
	streetName string
	city        string
	state       string
	customerId int
}

func getCustomer(db *sql.DB, id int) []Customer {
	query:="select * from customer inner join address on customer.ID=address.customer_ID where customer.ID"
	//sql enjection
	//sql placeholder
	var Ids[] interface{}
	if(id!=0){
		query="select * from customer inner join address on customer.ID=address.customer_ID where customer.ID=?"
		Ids=append(Ids,id)
	}
	rows,err:=db.Query(query,Ids...)
	if err != nil {
		panic(err.Error())
	}
	var result []Customer
	defer rows.Close()
	for rows.Next() {
		var cust Customer

		err = rows.Scan(&cust.id, &cust.name, &cust.DOB, &cust.add.id, &cust.add.streetName, &cust.add.city, &cust.add.state, &cust.add.customerId)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, cust)

	}
	return result
}
