// Copyright 2018-present the CoreDHCP Authors. All rights reserved
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"time"

	"github.com/coredhcp/coredhcp/config"
	"github.com/coredhcp/coredhcp/logger"
	"github.com/coredhcp/coredhcp/server"

	_ "github.com/coredhcp/coredhcp/plugins/dns"
	_ "github.com/coredhcp/coredhcp/plugins/file"
	_ "github.com/coredhcp/coredhcp/plugins/netmask"
	_ "github.com/coredhcp/coredhcp/plugins/range"
	_ "github.com/coredhcp/coredhcp/plugins/router"
	_ "github.com/coredhcp/coredhcp/plugins/server_id"
)

func main() {
	logger := logger.GetLogger("main")
	config, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}
	srv, err := server.Start(config)
	if err != nil {
		logger.Fatal(err)
	}
	if err := srv.Wait(); err != nil {
		logger.Print(err)
	}
	time.Sleep(time.Second)
}
