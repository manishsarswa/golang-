package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
type Customer struct{
	Name string
	Age int
	Address string
}
func handler(w http.ResponseWriter,r *http.Request){

	body, _ := ioutil.ReadAll(r.Body)
	var customer Customer
	err:=json.Unmarshal(body, &customer)

	if err!=nil{
		log.Fatal(err)
	}

	if customer.Age<18 {
		io.WriteString(w,"not eligible")
	}else {
		io.WriteString(w,string(body))
	}

}
func main() {
	http.HandleFunc("/customer/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
