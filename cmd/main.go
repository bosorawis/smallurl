package main

import (
	server "github.com/dihmuzikien/smallurl/url/port/http"
	"github.com/dihmuzikien/smallurl/url/repository/inmemory"
	"github.com/dihmuzikien/smallurl/url/usecase"
	"log"
	"net/http"
)

func main(){
	db := inmemory.New()
	urlSvc := usecase.NewUrlUsecase(db)
	srv, err := server.New(urlSvc)
	if err != nil {
		log.Fatalf("failed to initialize url server %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", srv))
}