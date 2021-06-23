package usecase

import (
	"context"
	"crypto/sha1"
	"github.com/dihmuzikien/smallurl/domain"
)

type urlUsecase struct {
	articleRepo    domain.UrlRepository
}

func (u *urlUsecase) Create(ctx context.Context, destination string) (domain.Url, error) {

	panic("implement me")
}

func (u *urlUsecase) CreateWithId(ctx context.Context, id, destination string) (domain.Url, error) {
	panic("implement me")
}

func (u *urlUsecase) GetById(ctx context.Context, id string) (domain.Url, error) {
	panic("implement me")
}

func (u *urlUsecase) List(ctx context.Context) ([]domain.Url, error) {
	panic("implement me")
}

func (u *urlUsecase) Delete(ctx context.Context, id string) error {
	panic("implement me")
}



func makeId(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	hashed := h.Sum(nil)
	return string(hashed)
}
