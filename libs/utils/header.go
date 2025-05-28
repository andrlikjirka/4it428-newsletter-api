package utils

import "net/http"

func GetXUserId(r *http.Request) string {
	return r.Header.Get("X-User-ID")
}
