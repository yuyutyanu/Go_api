package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"./pack"
)

func main(){
	res,err:= http.Get("http://localhost:18888")
	pack.ErrHandler(err)
	defer res.Body.Close()
	body,err:= ioutil.ReadAll(res.Body)
	pack.ErrHandler(err)
	log.Println("Status: ",res.Status)
	log.Println("StatusCode: ",res.StatusCode)
	log.Println("Headers: ",res.Header)
	log.Println(string(body))
}
