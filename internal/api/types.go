package api

import "github.com/gme-sh/gme.sh-api/pkg/gme-sh/short"

type SuccessableCreate struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    *short.ShortURL `json:"data"`
}

type SuccessablePool struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Pool    *short.Pool `json:"data"`
}
