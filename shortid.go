package service

import (
	"context"
	"github.com/teris-io/shortid"
)

func ShortId() (string, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return "", err
	}
	return sid.Generate()
}

func GenerateShortId(ctx context.Context) (string, error) {
	return ShortId()
}
