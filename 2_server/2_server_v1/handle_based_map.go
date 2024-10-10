package main

import (
	"net/http"
	"sync"
)

type Routable interface {
	Route(method string, pattern string, handler handlerFunc)
}

type Handler interface {
	ServeHTTP(context *Context)
	Routable
}

type HandlerBasedMap struct {
	//key 是method + url
	handlers sync.Map
}

func (h *HandlerBasedMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers.Load(key); ok {
		handler.(handlerFunc)(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("not any router match"))
		return
	}
}

// 一种常用的go设计模式，用于确保HandlerBasedMap实现了Handler接口
var _ Handler = &HandlerBasedMap{}

func (h *HandlerBasedMap) Route(method string, pattern string, handlerFunc handlerFunc) {
	key := h.key(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func (h *HandlerBasedMap) key(method string, path string) string {
	return method + "#" + path
}

func NewHandlerBasedMap() Handler {
	return &HandlerBasedMap{}
}
