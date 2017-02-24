package main

import (
	"go.uber.org/fx/dig"
	"go.uber.org/zap"
)

func main() {

	// start with a single container
	container := dig.New()

	container.Register(newLogger)
	container.Register(newClient)
	container.Register(newHandler)

	// resolve dispatcher
	var h *handler
	container.ResolveAll(&h)

	h.Hello()
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

type client struct {
	name string
}

func newClient() *client {
	return &client{name: "Barnaby"}
}

type handler struct {
	l *zap.Logger
	c *client
}

func (h *handler) Hello() {
	h.l.Info("Hello!", zap.Any("client", h.c))
}

func newHandler(l *zap.Logger, c *client) *handler {
	return &handler{l: l, c: c}
}
