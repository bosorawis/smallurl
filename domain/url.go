package domain

import (
	"context"
	"errors"
	"time"
)

type RepoError struct {
	Op 	string
	Err error
}



var (
	RepoPutError = errors.New("failed to save URL")
	RepoGetNotFoundError = errors.New("cannot find URL with matching ID")
	RepoListError = errors.New("failed to list URLs")
	RepoDeleteError = errors.New("failed to delete URL")
)

// Url represents model of URL redirection in the storage
type Url struct {
	ID string
	Destination string
	Created time.Time
}

type UrlRepository interface {
	Get(ctx context.Context, id string) (Url, error)
	Put(ctx context.Context, url Url) error
	List(ctx context.Context) ([]Url, error)
	Delete(ctx context.Context, id string) error
}

type UrlUseCase interface {
	Create(ctx context.Context, destination string) (Url, error)
	CreateWithId(ctx context.Context, id, destination string) (Url, error)
	GetById(ctx context.Context, id string) (Url, error)
	List(ctx context.Context) ([]Url, error)
	Delete(ctx context.Context, id string) error
}