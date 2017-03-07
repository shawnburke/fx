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

package uhttp

import (
	"context"

	"go.uber.org/fx/config"
	"go.uber.org/fx/modules/uhttp/client"
	"go.uber.org/fx/service"
	"go.uber.org/fx/ulog"
	"go.uber.org/zap"
)

const _moduleName = "http"

type Handlers struct {
	List []RouteHandler

	Inbound  []InboundMiddleware
	Outbound []client.OutboundMiddleware
}

// XNew is actually a DIG constructor...
func XNew(hs *Handlers, config config.Provider) service.ModuleProvider {
	return service.ModuleProviderFromFunc(_moduleName, func(host service.Host) (service.Module, error) {
		return newXModule(host, hs.List, config)
	})
}

func newXModule(
	_ service.Host, // NOTE: Host is now completely unused!!!
	hs []RouteHandler,
	config config.Provider,
) (*Module, error) {
	// setup config defaults
	cfg := Config{
		Port:    defaultPort,
		Timeout: defaultTimeout,
	}

	log := ulog.Logger(context.Background()).With(zap.String("module", _moduleName))

	if err := config.Scope("modules").Get(_moduleName).PopulateStruct(&cfg); err != nil {
		log.Error("Error loading http module configuration", zap.Error(err))
	}

	module := &Module{
		handlers: addHealth(hs),
		mcb:      defaultInboundMiddlewareChainBuilder(log, nil, nil),
		config:   cfg,
		log:      log,
	}

	//module.mcb = module.mcb.AddMiddleware(moduleOptions.inboundMiddleware...)

	return module, nil
}
