package main

import (
	"bytes"
	"mime/multipart"
	"./pack"
	"os"
	"io"
	"net/http"
	"log"
	"net/textproto"
)

func main(){
	var buffer bytes.Buffer
	writer:= multipart.NewWriter(&buffer)
	writer.WriteField("name","michael jackson")
	part := make(textproto.MIMEHeader)
	part.Set("Content-Type","image/jpeg")
	part.Set("Content-Disposition",`form-data; name="readme; filename=READEME.md"`)
	fileWriter,err := writer.CreatePart(part)
	pack.ErrHandler(err)
	readFile,err := os.Open("README.md")
	pack.ErrHandler(err)
	defer readFile.Close()
	io.Copy(fileWriter,readFile)
	writer.Close()

	res,err := http.Post("http://localhost:18888",writer.FormDataContentType(),&buffer)
	pack.ErrHandler(err)
	log.Println("Status:",res.Status)
}
