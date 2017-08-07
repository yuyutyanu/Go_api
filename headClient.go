package main

import (
	"net/http"
	"log"
)

func main() {
	res,err:=http.Head("http://localhost:18888")
	if err != nil{
		panic(err)
	}
	log.Println("Status: " ,res.Status)
	log.Println("Header: " ,res.Header)
	log.Println("Body: ",res.Body)//空
}
