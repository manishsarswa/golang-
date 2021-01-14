package main

import "database/sql"

type Customer struct{
	id int
	name string
	add Address
}
type Address struct{
	id int
	city string
	cid int
}

func handle(db *sql.DB,id int) []Customer{
	if id==0 {
		sq,err:=db.Query("select * from customer inner join address on customer.id=address.cid")
		if err!=nil{
			panic(err)
		}
		var c []Customer
		defer sq.Close()
		for sq.Next() {
			var cust Customer

			err = sq.Scan(&cust.id, &cust.name, &cust.add.id, &cust.add.city, &cust.add.cid)
			if err != nil {
				panic(err.Error())
			}
			c = append(c, cust)

		}
		return c


	}else{
		sq,err:=db.Query("select * from customer inner join address on customer.id=address.cid where customer.id=?",id)
		if err!=nil{
			panic(err)
		}
		var c []Customer
		defer sq.Close()
		for sq.Next() {
			var cust Customer

			err = sq.Scan(&cust.id, &cust.name, &cust.add.id, &cust.add.city, &cust.add.cid)
			if err != nil {
				panic(err.Error())
			}
			c = append(c, cust)

		}
		return c
	}
}
