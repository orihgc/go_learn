package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJson 读取请求的body，反序列化为obj
func (c *Context) ReadJson(obj interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(c.W, "read body failed: %v", err)
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		fmt.Fprintf(c.W, "deserialized body failed: %v", err)
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}
