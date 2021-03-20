package api

import (
	"errors"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"
	"github.com/gme-sh/gme.sh-api/pkg/gme-sh/shortreq"
	"github.com/imroc/req"
	"log"
)

func (a *API) GetPool(id *short.PoolID, secret string) (s *SuccessablePool, err error) {
	var res *req.Resp
	url := a.ApiUrl + "/pool/" + id.String() + "/" + secret
	log.Println("Requesting pool [name, secret] =", id.String(), "/", url)
	if res, err = req.Get(url); err != nil {
		return
	}
	log.Println("Resp:", res.String())
	err = res.ToJSON(&s)
	return
}

func (a *API) AppendPool(id *short.PoolID, name, secret, url string) (s *shortreq.Successable, err error) {
	var res *req.Resp
	if res, err = req.Post(a.ApiUrl+"/pool/"+id.String()+"/"+secret, req.BodyJSON(&shortreq.UpdatePoolPayload{
		Name: name,
		URL:  url,
	})); err != nil {
		return
	}
	err = res.ToJSON(&s)
	return
}

func (a *API) _CreateShortURL(p *shortreq.CreateShortURLPayload) (s *SuccessableCreate, err error) {
	res, err := a.post("/create", p)
	if err != nil {
		return nil, err
	}

	s = new(SuccessableCreate)
	err = res.ToJSON(s)

	return
}

func (a *API) _DeleteShortURL(url, secret string) (s *shortreq.Successable, err error) {
	if url = ExtractIDFromURL(url); url == "" {
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
