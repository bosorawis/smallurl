package goapp

import "time"

type Repository interface {
	Get(id string) (string, error)
	Put(url string) (*Url, error)
}

type Url struct {
	ID string
	DestinationURL string
	Created time.Time
	Expired time.Time
}


