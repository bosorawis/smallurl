package inmemory

import (
	"context"
	"github.com/dihmuzikien/smallurl/domain"
	"time"
)

type Db struct {
	storage map[string]data
}

type data struct {
	dest string
	created time.Time
}

func New() *Db{
	return &Db{
		storage: make(map[string]data),
	}

}
func (d *Db) Put(ctx context.Context, url domain.Url) error {
	d.storage[url.ID] = data {
		dest: url.Destination,
		created: url.Created,
	}
	return nil
}

func (d *Db) Delete(ctx context.Context, id string) error {
	delete(d.storage, id)
	return nil
}

func (d *Db) Get(ctx context.Context, id string) (domain.Url, error) {
	if _, ok := d.storage[id]; !ok {
		return domain.Url{}, domain.ErrNotFound
	}
	u := d.storage[id]
	return domain.Url{
		ID: id,
		Destination: u.dest,
		Created: u.created,
	}, nil
}

func (d *Db) List(ctx context.Context) ([]domain.Url, error) {
	var urls []domain.Url
	for k, v := range d.storage{
		urls = append(urls, domain.Url{ID: k, Destination: v.dest, Created: v.created})
	}
	return urls, nil
}


