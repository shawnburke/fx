// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package xservice

import (
	"log"

	"go.uber.org/fx/config"
	"go.uber.org/fx/dig"
	"go.uber.org/fx/service"
	"go.uber.org/zap"
)

type Builder struct {
	g *dig.Graph
}

func WithModule(m interface{}, constructors ...interface{}) *Builder {
	g := dig.New()
	b := &Builder{g: g}
	return b.WithModule(m, constructors...)
}

func (b *Builder) WithModule(m interface{}, constructors ...interface{}) *Builder {
	b.g.MustRegister(m)

	// Register all the provided module specific constructors
	for _, c := range constructors {
		b.g.MustRegister(c)
	}

	return b
}

func (b *Builder) Build() service.Manager {
	// Sample logger object inserted in for demo purposes
	l, _ := zap.NewProduction()
	b.g.MustRegister(l)

	// Lets get config constructor going in the graph, so anyone who needs it can get it
	b.g.MustRegister(config.Load)

	// Resolve ModuleProvider from the graph
	// In the future this can be "find all ModuleProviders in graph
	var mp service.ModuleProvider
	b.g.MustResolve(&mp)

	m, err := service.WithModule(mp).Build()
	if err != nil {
		log.Panic(err)
	}

	return m
}
