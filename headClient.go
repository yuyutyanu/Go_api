package main

import (
	"net/http"
	"log"
	"./pack"
)

func main() {
	res,err:=http.Head("http://localhost:18888")
	pack.ErrHandler(err)
	log.Println("Status: " ,res.Status)
	log.Println("Header: " ,res.Header)
	log.Println("Body: ",res.Body)//空
}
