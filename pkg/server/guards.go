package server

import (
	"net/http"
	"strings"
)

func LoginGuard(r *http.Request) bool {
	return true
}

func PostExistsGuard(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func MethodsGuard(w http.ResponseWriter, r *http.Request, methods ...string) bool {
	method := strings.ToUpper(r.Method)

	for _, v := range methods {
		if method == strings.ToUpper(v) {
			return true
		}
	}

	return false
}
