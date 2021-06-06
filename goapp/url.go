package goapp

import (
	"time"
)

type Repository interface {
	Get(id string) (string, error)
	Put(url, id string) error
}

type Url struct {
	ID string
	DestinationURL string
	Created time.Time
	Expired time.Time
}

func Put(r Repository, url string) (string, error){

	return "", nil
}



