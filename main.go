package main

import (
	"fmt"
	"log"

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

// DomainsCTRL ...
func DomainsCTRL(ctx *fasthttp.RequestCtx) {
	h, ok := domains[string(ctx.Host())]
	if !ok {
		ctx.NotFound()
		return
	}
	h(ctx)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", DomainsCTRL)
	router.POST("/listener", ListenerCTRL)
	router.DELETE("/listener/:name", ListenerUnregisterCTRL)
	router.POST("/publish/{event}", PublishEventCTRL)
	server, err := registerHost(router.Handler)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(server.ListenAndServeTLS(":8080", "", ""))
}
