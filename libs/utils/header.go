package utils

import (
	"net/http"
)

func GetXUserId(r *http.Request) string {
	return r.Header.Get("X-User-ID")
}

func GetNewsletterIdFromQueryParam(r *http.Request) string {
	return r.URL.Query().Get("newsletter_id")
}
