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
	handler Handler
	root    Filter
}

// Route 实现 Server 接口,注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(context *Context)) {
	s.handler.Route(method, pattern, handlerFunc)
}

// Start 实现 Server 接口，
func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
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

func NewHttpServer(name string, filters ...FilterBuilder) Server {
	handler := NewHandlerBasedOnTree()
	//handler := NewHandlerBasedMap()
	var root Filter = handler.ServeHTTP
	for i := len(filters) - 1; i >= 0; i-- {
		builder := filters[i]
		root = builder(root)
	}
	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

func main() {
	server := NewHttpServer("test-server", MetricsFilterBuilder)
	server.Route(http.MethodPost, "/signup", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
		return
	}
}
