package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, r *http.Request) {
	b := []byte("Hello world")
fmt.Println(strings.Replace(r.RequestURI,"/","",-1))
	w.Write(b)
	s:=r.RequestURI
	s=strings.Replace(s,"/","",-1)

	io.WriteString(w, s)
}
func Printer(w http.ResponseWriter, r *http.Request) {
	b := []byte("Printing ")
	w.Write(b)
	s:=r.RequestURI
	var x []string=strings.Split(s, "/")
	io.WriteString(w,x[2])
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/print/", Printer)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

