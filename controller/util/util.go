package util

import (
	"net/http"
	"strings"
)

func HasSubURI(r *http.Request) (bool, string) {
	url := strings.Split(r.URL.Path, "/")
	return (len(url[2]) > 0), url[2]
}
