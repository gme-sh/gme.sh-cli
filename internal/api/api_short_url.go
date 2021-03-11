package api

import (
	"errors"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
)

func (a *API) CreateShortURL(p *shortreq.CreateShortURLPayload) (s *SuccessableCreate, err error) {
	res, err := a.post("/create", p)
	if err != nil {
		return nil, err
	}

	s = new(SuccessableCreate)
	err = res.ToJSON(s)

	return
}

func (a *API) DeleteShortURL(url, secret string) (s *shortreq.Successable, err error) {
	if url = extractIDFromURL(url); url == "" {
		return nil, errors.New("no url given")
	}
	res, err := a.delete("/" + url + "/" + secret)
	if err != nil {
		return nil, err
	}
	s = new(shortreq.Successable)
	err = res.ToJSON(s)
	return
}
