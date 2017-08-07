package main

import (
	"net/url"
	"net/http"
	"log"
	//"io/ioutil"
	"io/ioutil"
)

func main(){
	values:= url.Values{
		"query":{"hello world"},
	}
	res,_:=http.Get("http://localhost:18888?"+values.Encode())
	defer res.Body.Close()
	log.Println(values.Encode())
	Body,_ :=ioutil.ReadAll(res.Body)
	log.Println(string(Body))
}
