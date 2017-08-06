package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"helloworld")
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":18888",nil)
}
