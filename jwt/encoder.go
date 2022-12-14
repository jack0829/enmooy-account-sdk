package jwt

import "github.com/speps/go-hashids/v2"

func NewEncoder(salt string) *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 16
	h, _ := hashids.NewWithData(hd)
	return h
}
