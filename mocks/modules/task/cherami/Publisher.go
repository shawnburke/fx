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

package cherami

import cherami "github.com/uber/cherami-client-go/client/cherami"
import mock "github.com/stretchr/testify/mock"

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Publisher) Close() {
	_m.Called()
}

// Open provides a mock function with given fields:
func (_m *Publisher) Open() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: message
func (_m *Publisher) Publish(message *cherami.PublisherMessage) *cherami.PublisherReceipt {
	ret := _m.Called(message)

	var r0 *cherami.PublisherReceipt
	if rf, ok := ret.Get(0).(func(*cherami.PublisherMessage) *cherami.PublisherReceipt); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cherami.PublisherReceipt)
		}
	}

	return r0
}

// PublishAsync provides a mock function with given fields: message, done
func (_m *Publisher) PublishAsync(message *cherami.PublisherMessage, done chan<- *cherami.PublisherReceipt) (string, error) {
	ret := _m.Called(message, done)

	var r0 string
	if rf, ok := ret.Get(0).(func(*cherami.PublisherMessage, chan<- *cherami.PublisherReceipt) string); ok {
		r0 = rf(message, done)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cherami.PublisherMessage, chan<- *cherami.PublisherReceipt) error); ok {
		r1 = rf(message, done)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ cherami.Publisher = (*Publisher)(nil)
