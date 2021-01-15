package TestSql


import (
	"database/sql"
	"strconv"

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

func getCustomer(db *sql.DB, id string) []Customer {
	query:="select * from customer inner join address on customer.ID=address.customer_ID where customer.ID"
	//sql enjection
	//sql placeholder
	var Ids[] interface{}
	if(id!="0"){
		query="select * from customer inner join address on customer.ID=address.customer_ID where customer.ID=?"
		x,err:=strconv.Atoi(id)
		if err!=nil{

		}
		Ids=append(Ids,x)
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
//func Handler(w http.ResponseWriter,r *http.Request){
//	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//	if err!=nil{
//		panic(err)
//	}
//	s:=r.RequestURI
//	var x []string=strings.Split(s, "/")
//	xx,err:=strconv.Atoi(x[1])
//	rows,err:=db.Query("select * from customer inner join address on customer.ID=address.customer_ID where customer.ID=?", xx)
//	var result []Customer
//	for rows.Next(){
//		var cust Customer
//		err = rows.Scan(&cust.id, &cust.name, &cust.DOB, &cust.add.id, &cust.add.streetName, &cust.add.city, &cust.add.state, &cust.add.customerId)
//		if err != nil {
//			panic(err.Error())
//		}
//		//ss=ss+string(cust.ID)
//		io.WriteString(w,strings.Itoi(&cust.id))
//
//
//		result = append(result, cust)
//	}
//
//}
//func main(){
//	http.HandleFunc("/",Handler)
//	log.Fatal("8080",nil)
//
//}
