package main

import (
	"github.com/dihmuzikien/smallurl/goapp/server"
	"github.com/dihmuzikien/smallurl/goapp/storage/inmemory"
	"log"
)

func main(){
	db := inmemory.NewDb()
	server,_ := server.NewServer(db)
	log.Fatal(server.Start(":8080"))
}