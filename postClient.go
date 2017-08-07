package main

import (
	"os"
	"net/http"
	"./pack"
	"log"
)

func main() {
	file,err:= os.Open("postClient.go")
	pack.ErrHandler(err)
	res,err:=http.Post("http://localhost:18888","text/plain",file)
	pack.ErrHandler(err)
	log.Println("Status: ",res.Status)
}
