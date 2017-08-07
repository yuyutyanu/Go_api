package main

import (
	"net/http"
	"net/url"
	"log"
)

func main() {
	values:= url.Values{
		"query":{"helloworld"},
	}
	res,err:=http.PostForm("http://localhost:18888",values)
	if err != nil{
		panic(err)
	}
	log.Println(res.Status)
}
