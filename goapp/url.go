package goapp

import "crypto/sha1"

type Repository interface {
	Get(id string) (string, error)
	Put(url, id string) error
}

func Put(r Repository, url string) (string, error){

	return "", nil
}

func makeId(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	hashed := h.Sum(nil)
	return string(hashed)
}
