package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Customer struct{
	Id int
	Name string
	Dob string
	Address string
}

//
//func handleGet(w http.ResponseWriter,r *http.Request){
//
//	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//	if err!=nil{
//		panic(err)
//	}
//
//	var idr[] interface {}
//	//vars:=mux.Vars(r)
//	Name:=r.URL.Query()["name"][0]
//	idr=append(idr,Name)
//
//	rows,err:=db.Query(`select * from cust where name=?`,idr...)
//	if err!=nil {
//		panic(err)
//	}
//
//	var result []Customer
//
//	defer rows.Close()
//	for rows.Next() {
//		var cust Customer
//
//		err = rows.Scan(&cust.Id, &cust.Name,&cust.Dob,&cust.Address)
//		result = append(result, cust)
//
//		fmt.Println(result)
//
//	}
//	json.NewEncoder(w).Encode(result)
//
//
//
//}
//
//
//func handleGetid(w http.ResponseWriter,r *http.Request){
//	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//	if err!=nil{
//
//	}
//
//	defer db.Close()
//	vars:=mux.Vars(r)
//	//fmt.Println("response is", r)
//	id:=vars["id"]
//	fmt.Println("id is=",id)
//
//	query:="select * from cust"
//	var ids[] interface{}
//	if(id!="0"){
//		query=`select * from cust where id=?`
//		x,_:=strconv.Atoi(id)
//		ids=append(ids,x)
//	}
//	rows,_:=db.Query(query,ids...)
//	defer rows.Close()
//	var result []Customer
//
//	for rows.Next(){
//		var c Customer
//		rows.Scan(&c.Id,&c.Name,&c.Dob,&c.Address)
//		//fmt.Println(c)
//		result=append(result,c)
//	}
//	//fmt.Println("i func ", c)
//
//	json.NewEncoder(w).Encode(result)
//
//}
//
//
//func handlePost(w http.ResponseWriter,r *http.Request){
//	db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//		if err!=nil{
//			fmt.Print(err)
//		}
//	var idr[] interface{}
//	var c Customer
//	query:=`insert into cust value(?,?,?,?)`
//	body,_:=ioutil.ReadAll(r.Body)
//	json.Unmarshal(body,&c)
//	idr=append(idr,&c.Id)
//	idr=append(idr,&c.Name)
//	idr=append(idr,&c.Dob)
//	idr=append(idr,&c.Address)
//	id:=&c.Id
//	fmt.Println("customer details is", c)
//	_,err=db.Exec(query,idr...)
//	if err!=nil{
//		fmt.Print(err)
//	}
//
//	query=`select * from cust where id=?`
//	var idd[] interface{}
//	idd=append(idd,id)
//	rows,err:=db.Query(query,idd...)
//	var result []Customer
//	for rows.Next(){
//		var c Customer
//		rows.Scan(&c.Id,&c.Name,&c.Dob,&c.Address)
//		result=append(result,c)
//	}
//
//	json.NewEncoder(w).Encode(result)
//
//
//
//}



func updateCustomer(w http.ResponseWriter,r *http.Request){
		db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
		if err!=nil{
			fmt.Print(err)
		}
		vars:=mux.Vars(r)
		ide:=vars["id"]
		query:=`update cust set`
		var idr [] interface{}

		var c Customer
		body,_:=ioutil.ReadAll(r.Body)
		json.Unmarshal(body,&c)

		if c.Name!=""{
			query+=" name=?"
			idr=append(idr,c.Name)
		}
		if c.Address!=""{
			query+=", address=?"
			idr=append(idr,c.Address)
		}
		query+=" where id=?"
		fmt.Println(query)
		idr=append(idr,ide)
		_,er:=db.Exec(query,idr...)

		if er!=nil{
			fmt.Print(er)
		}

		query=`select * from cust where id=?`
		var idd[] interface{}
		idd=append(idd,ide)
		rows,_:=db.Query(query,idd...)
		var result []Customer
		for rows.Next(){
			var c Customer
			rows.Scan(&c.Id,&c.Name,&c.Dob,&c.Address)
			result=append(result,c)
			fmt.Println(c)
			fmt.Println(c)
		}
		json.NewEncoder(w).Encode(result)

}

//
//
//func handleDelete(w http.ResponseWriter,r *http.Request){
//		db,err:=sql.Open("mysql","root:Manish@123Sharma@/Customer_service")
//		defer db.Close()
//		if err!=nil{
//			fmt.Print(err)
//		}
//
//		vars:=mux.Vars(r)
//	    id,_:=strconv.Atoi(vars["id"])
//	    var idr[] interface{}
//	    idr=append(idr,id)
//		var result []Customer
//		query:=`select * from cust where id=?`
//		rows,_:=db.Query(query,idr...)
//		for rows.Next(){
//			var c Customer
//			rows.Scan(&c.Id,&c.Name,&c.Dob,&c.Address)
//			result=append(result,c)
//		}
//		query=`delete from cust where id=?`
//		_,err=db.Exec(query,idr...)
//		json.NewEncoder(w).Encode(result)
//
//}


func main(){
	r:=mux.NewRouter()
	//r.HandleFunc("/customer",handleGet).Methods(http.MethodGet)
	//r.HandleFunc("/customer/{id}",handleGetid).Methods(http.MethodGet)
	//r.HandleFunc("/customer/",handlePost).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}",updateCustomer).Methods(http.MethodPut)
	//r.HandleFunc("/customer/{id}",handleDelete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":3003",r))
}
