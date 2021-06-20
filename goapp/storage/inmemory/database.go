package inmemory

import (
	"context"
	"github.com/dihmuzikien/smallurl/goapp/usecase/url"
)

type Db struct {
	storage map[string]string
}

func NewDb() (*Db){
	return &Db{
		storage: make(map[string]string),
	}

}

func (d *Db) Get(ctx context.Context, id string) (string, error) {
	if _, ok := d.storage[id]; !ok {
		return "", url.ErrNotFound
	}
	return d.storage[id], nil
}

func (d *Db) Put(ctx context.Context, id, url string) error {
	d.storage[id] = url
	return nil
}

func (d *Db) List(ctx context.Context) ([]url.RedirectModel, error) {
	var urls []url.RedirectModel
	for k, v := range d.storage{
		urls = append(urls, url.RedirectModel{ID: k, Destination: v})
	}
	return urls, nil
}


