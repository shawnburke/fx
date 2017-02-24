package main

import (
	"go.uber.org/fx/dig"
	"go.uber.org/zap"
)

// NewService creates an app framework service
func NewService() *Service {
	return &Service{container: dig.New()}
}

// Service is the service being bootstrapped
type Service struct {
	container dig.Graph
	types     []interface{}
}

// RegisterType adds a new
func (s *Service) RegisterType(t interface{}) {
	s.types = append(s.types, t)
}

// Start builds the container, registers handlers
// and starts the messaging framework
func (s *Service) Start() {

	// register framework types
	s.container.Register(newLogger)

	// register user types
	for _, t := range s.types {
		s.container.Register(t)
	}

	// handlers should be registered with RegisterHandler
	// which we can resolve and register with the dispatcher
	// before calling dispatcher.Start()
	var h *handler
	s.container.ResolveAll(&h)

	h.Hello()
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
