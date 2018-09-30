package main

import (
	"fmt"

	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// ListenerRequestDTO ...
type ListenerRequestDTO struct {
	Event   string `json:"event" valid:"required"`
	Name    string `json:"name" valid:"required"`
	Address string `json:"address" valid:"required"`
}

// ListenerCTRL ...
func ListenerCTRL(ctx *fasthttp.RequestCtx) {
	var listener ListenerRequestDTO
	ctx.Request.BodyWriter()
	err := json.Unmarshal(ctx.PostBody(), &listener)
	if err != nil {

	}
	_, err = govalidator.ValidateStruct(listener)
}

// ListenerUnregisterCTRL ...
func ListenerUnregisterCTRL(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

// PublishEventCTRL ...
func PublishEventCTRL(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	router := fasthttprouter.New()
	router.POST("/listener", ListenerCTRL)
	router.DELETE("/listener/:name", ListenerUnregisterCTRL)
	router.POST("/publish/{event}", PublishEventCTRL)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
