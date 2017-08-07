package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter,r *http.Request){
	req,err:=httputil.DumpRequest(r,true)
	if err !=nil{
		panic(nil)
	}
	fmt.Println(string(req))
	fmt.Fprintf(w,"helloworld")
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":18888",nil)
}
