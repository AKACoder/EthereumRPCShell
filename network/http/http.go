/*
 * Copyright 2020 The SealABC Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package http

import (
	"errors"
	"github.com/AKACoder/EthereumRPCShell/common/wLog"
	"github.com/gin-gonic/gin"
	"sync"
)

type Server struct {
	Config *Config
}

func (s *Server) Start(wg *sync.WaitGroup) (*sync.WaitGroup, error) {
	var err error
	if s.Config == nil {
		err = errors.New("no configure")
		return nil, err
	}

	router := gin.Default()

	//using before middlewares
	router.Use(middlewares.before...)

	s.setRoutes(router, *s.Config)

	if wg == nil {
		wg = &sync.WaitGroup{}
	}
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				wLog.Log.Error(e)
			}
			wg.Done()
		}()

		err = router.Run(s.Config.Address)
		if err != nil {
			wLog.Log.Warn("start http server failed: ", err.Error())
		}
	}()

	return wg, err
}

func (s *Server) setRoutes(router gin.IRouter, cfg Config) {
	for _, regTo := range cfg.RouteRegisters {
		regTo(router)
	}
}
