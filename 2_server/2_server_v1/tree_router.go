package main

import (
	"net/http"
	"strings"
)

type HandlerBaseOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func (h HandlerBaseOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h HandlerBaseOnTree) findRouter(path string) (handlerFunc, bool) {
	// 去除头尾可能有的/，然后按照/切割成段
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前 path 的节点
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	// 到这里，应该是找完了
	if cur.handler == nil {
		// 到达这里是因为这种场景
		// 比如说你注册了 /user/profile
		// 然后你访问 /user
		return nil, false
	}
	return cur.handler, true
}

func (h HandlerBaseOnTree) Route(method string, pattern string, handler handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root
	for index, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handler)
		}
	}
}

func (h HandlerBaseOnTree) createSubTree(root *node, paths []string, handler handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handler
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

func (h HandlerBaseOnTree) findMatchChild(n *node, path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBaseOnTree{
		root: &node{},
	}
}
