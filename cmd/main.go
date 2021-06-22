package main

import (
	"github.com/dihmuzikien/smallurl/goapp/server"
	"github.com/dihmuzikien/smallurl/goapp/storage/inmemory"
	"log"
	"net/http"
)

func main(){
	db := inmemory.NewDb()
	server,_ := server.NewServer(db)
	log.Fatalln(http.ListenAndServe(":8080", server))
}