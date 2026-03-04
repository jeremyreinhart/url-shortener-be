package utils

import "net/url"

func IsValidURL(input string) bool {
	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	return u.Scheme == "http"  || u.Scheme == "https"
}