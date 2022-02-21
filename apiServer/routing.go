package apiServer

import (
	"io"
	"net/http"
)

func (s *ApiServer) HandleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}