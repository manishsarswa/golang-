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

	var stmtOut *sql.Rows
	var err error

	if id == 0 {
		stmtOut, err = db.Query("SELECT * FROM customer INNER JOIN address ON customer.ID = address.customer_ID;")
	} else {
		stmtOut, err = db.Query("select * from customer inner join address on customer.ID=address.customer_ID where customer.ID=?", id)
	}

	if err != nil {
		panic(err.Error())
	}
	var result []Customer
	defer stmtOut.Close()
	for stmtOut.Next() {
		var cust Customer

		err = stmtOut.Scan(&cust.id, &cust.name, &cust.DOB, &cust.add.id, &cust.add.streetName, &cust.add.city, &cust.add.state, &cust.add.customerId)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, cust)

	}
	return result
}
