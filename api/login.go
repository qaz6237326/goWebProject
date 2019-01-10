package api

import (
	"net/http"
)

func init() {
	http.HandleFunc("/login", Login)
}
