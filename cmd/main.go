package main

import (
	http2 "github.com/dihmuzikien/smallurl/internal/port/http"
	inmemory2 "github.com/dihmuzikien/smallurl/internal/storage/inmemory"
	usecase2 "github.com/dihmuzikien/smallurl/internal/usecase"
	"log"
	"net/http"
)

func main() {
	db := inmemory2.New()
	urlSvc := usecase2.NewUrlUseCase(db)
	srv := http2.New(urlSvc)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
