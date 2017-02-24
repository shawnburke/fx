package main

import "go.uber.org/zap"

func main() {
	service := NewService()

	service.RegisterType(newClient)
	service.RegisterType(newHandler)

	service.Start()
}

type client struct {
	name string
}

func newClient() *client {
	return &client{name: "Highgarden"}
}

type handler struct {
	l *zap.Logger
	c *client
}

func (h *handler) Hello() {
	h.l.Info("Hello!", zap.Any("client", h.c.name))
}

func newHandler(l *zap.Logger, c *client) *handler {
	return &handler{l: l, c: c}
}
