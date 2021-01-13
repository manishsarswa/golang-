package TestSql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

type Customer struct{
	ID int
	name string
	DOB string
	add address
}
type address struct{
	ID int
	street_Name string
	city string
	state string
	customer_ID int
}
//func handle(w http.ResponseWriter, r *http.Request){
//
//	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//	if err!=nil{
//		panic(err)
//	}
//
//	defer db.Close()
//	sq,err:=db.Query("select * from customer where ID=3")
//	//stmtIns, err := db.Prepare("INSERT INTO customer VALUES( ?, ?, ?)") // ? = placeholder
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//	defer sq.Close()
//	for sq.Next(){
//				var cust Customer
//
//				err=sq.Scan(&cust.ID,&cust.name,&cust.DOB)
//				if err!=nil{
//					panic(err.Error())
//				}
//
//				io.WriteString(w,string(cust.ID))
//				io.WriteString(w,cust.name)
//		io.WriteString(w,cust.DOB)
//
//
//
//	}
//	//for i := 5; i < 25; i++ {
//	//	_, err = sq.Exec(i, "Manish Sharma", "15-02-1998") // Insert tuples (i, i^2)
//	//	if err != nil {
//	//		panic(err.Error()) // proper error handling instead of panic in your app
//	//	}
//	//}
//}
func getcustomer(db *sql.DB,id int) []Customer{
//	http.HandleFunc("/",handle)
//	log.Fatal(http.ListenAndServe(":8087", nil))

	if id == 0 {
		stmtOut, err := db.Query("SELECT * FROM customer INNER JOIN address ON customer.ID = address.customer_ID;")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtOut.Close()
		var res []Customer
		for stmtOut.Next() {

			var c Customer
			if err := stmtOut.Scan(&c.ID, &c.name, &c.DOB, &c.add.ID, &c.add.street_Name, &c.add.city, &c.add.state, &c.add.customer_ID); err != nil {
				log.Fatal(err)
			}
			//fmt.Println(c.ID, c.Name, c.DOB, c.Addr.ID, c.Addr.StreetName, c.Addr.City, c.Addr.State, c.Addr.Cus_ID)
			res = append(res, c)
		}

		return res
	}else {
		sq, err := db.Query("select * from customer inner join address on customer.ID=address.customer_ID where customer.ID=?", id)
		//stmtIns, err := db.Prepare("INSERT INTO customer VALUES( ?, ?, ?)") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var c []Customer
		defer sq.Close()
		for sq.Next() {
			var cust Customer

			err = sq.Scan(&cust.ID, &cust.name, &cust.DOB, &cust.add.ID, &cust.add.street_Name, &cust.add.city, &cust.add.state, &cust.add.customer_ID)
			if err != nil {
				panic(err.Error())
			}
			//ss=ss+string(cust.ID)

			c = append(c, cust)

		}
		return c
	}
	//for i := 5; i < 25; i++ {
	//	_, err = sq.Exec(i, "Manish Sharma", "15-02-1998") // Insert tuples (i, i^2)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	//}
}



//package main
//
//import (
//"database/sql"
//"fmt"
//_ "github.com/go-sql-driver/mysql"
//)
//
//type Customer struct {
//	ID int `json:"ID"`
//	Name string `json:"Cust_Name"`
//	DOB string`json:"DOB"`
//}
//
//func main (){
//
//
//	db,err:=sql.Open("mysql","root:Password@/Customer_services")
//	if err!=nil {
//		panic(err.Error())
//	}
//
//	defer db.Close()
//
//	result,err:=db.Query("SELECT * from Customers")
//	if err!=nil{
//		panic(err.Error())
//	}
//
//
//	for result.Next(){
//		var cust Customer
//
//		err=result.Scan(&cust.ID,&cust.Name,&cust.DOB)
//		if err!=nil{
//			panic(err.Error())
//		}
//
//
//		fmt.Println(cust.ID, cust.Name, cust.DOB)
//
//	}
//
//
//}
