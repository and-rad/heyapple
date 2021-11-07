package mock

import "net/http"

type Handler struct {
	Body string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Body))
}
