package main

import (
	"net/http"
	"net/url"
	"log"
	"./pack"
)

func main() {
	values:= url.Values{
		"query":{"helloworld"},
	}
	res,err:=http.PostForm("http://localhost:18888",values)
	pack.ErrHandler(err)
	log.Println(res.Status)
}
