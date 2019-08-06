package middlewares

import (
    "http"
)

func logger(rw http.ResponseWriter, r *http.Request, next http.HandleFunc) {
    next(rw,r)
}