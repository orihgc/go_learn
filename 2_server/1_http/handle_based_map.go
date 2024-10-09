package main

import "net/http"

type HandlerBasedMap struct {
	//key 是method + url
	handlers map[string]func(context *Context)
}

func (h *HandlerBasedMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(w, r))
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not any router match"))
		return
	}
}

func (h *HandlerBasedMap) key(method string, path string) string {
	return method + "#" + path
}
