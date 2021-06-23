package domain

import (
	"context"
	"time"
)

// Url represents model of URL redirection in the repository
type Url struct {
	ID string
	Destination string
	Created time.Time
}

type UrlRepository interface {
	Get(ctx context.Context, id string) (string, error)
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