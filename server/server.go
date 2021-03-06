// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package server

import (
	"sync/atomic"

	"github.com/berkaroad/saashard/admin"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/proxy"
)

const (
	// Offline = 0
	Offline = iota
	// Online = 1
	Online
	// Unknown = 2
	Unknown
)

// Server is startup endpoint.
type Server struct {
	cfg *config.Config

	statusIndex int32
	status      [2]int32
	running     bool

	proxy *proxy.Server
	admin *admin.Server
}

// NewServer is to create a server.
func NewServer(cfg *config.Config) (*Server, error) {
	var err error
	s := new(Server)
	s.cfg = cfg

	atomic.StoreInt32(&s.statusIndex, 0)
	s.status[s.statusIndex] = Online

	s.proxy, err = proxy.NewServer(cfg)
	s.admin, err = admin.NewServer(cfg)
	return s, err
}

// Run server.
func (s *Server) Run() {
	s.running = true

	// proxy
	go s.proxy.Run()

	// admin
	s.admin.Run()
}

// Status of server.
func (s *Server) Status() string {
	var status string
	switch s.status[s.statusIndex] {
	case Online:
		status = "online"
	case Offline:
		status = "offline"
	case Unknown:
		status = "unknown"
	default:
		status = "unknown"
	}
	return status
}

// Close server.
func (s *Server) Close() {
	s.running = false
	s.proxy.Close()
	s.admin.Close()
}
