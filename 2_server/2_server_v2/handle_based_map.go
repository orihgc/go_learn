package main

import "net/http"

type Routable interface {
	Route(method string, pattern string, handler func(context *Context))
}

type Handler interface {
	http.Handler
	Routable
}

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

// 一种常用的go设计模式，用于确保HandlerBasedMap实现了Handler接口
var _ Handler = &HandlerBasedMap{}

func (h *HandlerBasedMap) Route(method string, pattern string, handlerFunc func(context *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

func (h *HandlerBasedMap) key(method string, path string) string {
	return method + "#" + path
}

func NewHandlerBasedMap() Handler {
	return &HandlerBasedMap{
		handlers: make(map[string]func(context *Context)),
	}
}
