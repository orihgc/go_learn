package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(context *Context))

	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBasedMap
}

// Route 实现 Server 接口,注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(context *Context)) {
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handlerFunc
}

// Start 实现 Server 接口，
func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func SignUp(context *Context) {
	signUpReq := &signUpRequest{}
	err := context.ReadJson(signUpReq)
	if err != nil {
		fmt.Println(err)
		err = context.BadRequestJson(err)
		return
	}
	resp := &signUpResponse{
		ID: 123,
	}

	err = context.OkJson(resp)
	if err != nil {
		err = context.BadRequestJson(err)
		return
	}

}

type signUpRequest struct {
	// 这里的tag 是json序列化的时候用的
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpResponse struct {
	ID int64 `json:"id"`
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
		handler: &HandlerBasedMap{
			handlers: make(map[string]func(context *Context)),
		},
	}
}

func main() {
	server := NewHttpServer("test-server")
	server.Route("POST", "/signup", SignUp)
	server.Start(":8080")
}
