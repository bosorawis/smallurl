package main

import (
	http2 "github.com/dihmuzikien/smallurl/url/port/http"
	"github.com/dihmuzikien/smallurl/url/repository/inmemory"
	"log"
	"net/http"
)

func main(){
	db := inmemory.New()
	server,_ := http2.NewServer(db)
	log.Fatalln(http.ListenAndServe(":8080", server))
}