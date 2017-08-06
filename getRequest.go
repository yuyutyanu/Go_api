package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func main(){
	res,err:= http.Get("http://localhost:18888")
	errHandler(err)
	defer res.Body.Close()
	body,err:= ioutil.ReadAll(res.Body)
	errHandler(err)
	log.Println("Status: ",res.Status)
	log.Println("StatusCode: ",res.StatusCode)
	log.Println("Headers: ",res.Header)
	log.Println(string(body))
}

func errHandler(err error){
	if err !=nil{
		panic(nil)
	}
}