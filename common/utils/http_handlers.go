package utils

import (
	"net/http"
)

func DelegatingHandler(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	next(rw, req)
}
