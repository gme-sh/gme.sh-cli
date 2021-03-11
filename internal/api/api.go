package api

import (
	"github.com/full-stack-gods/gme.sh-cli/internal/config"
	"github.com/imroc/req"
	"strings"
)

type API struct {
	ApiUrl string
}

func NewApi(cfg *config.Config) *API {
	return &API{ApiUrl: cfg.ApiUrl}
}

func (a *API) post(path string, payload interface{}) (res *req.Resp, err error) {
	res, err = req.Post(a.ApiUrl+path, req.BodyJSON(payload))
	return
}

func (a *API) delete(path string) (res *req.Resp, err error) {
	res, err = req.Delete(a.ApiUrl + path)
	return
}

func (a *API) GetURL(id string) string {
	return a.ApiUrl + "/" + id
}

func ExtractIDFromURL(url string) string {
	// extract id
	for strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	if strings.Contains(url, "/") {
		url = url[strings.LastIndex(url, "/")+1:]
	}
	return url
}
