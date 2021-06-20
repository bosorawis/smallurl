package url

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
)
var (
	ErrNotFound = errors.New("cannot find URL with matching key")
)

type Repository interface {
	Get(ctx context.Context, id string) (string, error)
	Put(ctx context.Context, id, url string) error
	List(ctx context.Context) ([]RedirectModel, error)
}

type RedirectModel struct {
	ID string
	Destination string
}

type Store struct {
	storage Repository
}

func NewStore(r Repository) (*Store, error){
	return &Store{
		storage: r,
	}, nil
}

func (s *Store) Get(ctx context.Context, id string) (string, error){
	url, err := s.storage.Get(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve url %w", err)
	}
	return url, nil
}

func (s *Store) PutWithAlias(ctx context.Context, alias, url string) error{
	err := s.storage.Put(ctx, alias, url)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store)  Put(ctx context.Context, url string) (string, error){
	id := makeId(url)
	err := s.storage.Put(ctx, id, url)
	if err != nil {
		return "", err
	}
	return id, nil
}

func makeId(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	hashed := h.Sum(nil)
	return string(hashed)
}
