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

	"go.uber.org/fx/dig"
	"go.uber.org/fx/service"
	"go.uber.org/fx/x/modules/xhttp"
)

type Service struct {
	service.Manager

	g    dig.Graph
	mods []service.Module
}

func New(cs ...interface{}) *Service {
	g := dig.New()

	// register all the provided constructors
	for _, c := range cs {
		g.MustRegister(c)
	}

	m, err := service.WithModule("http", xhttp.New(g)).Build()
	if err != nil {
		log.Panic(err)
	}

	return &Service{
		g:       g,
		Manager: m,
	}
}
