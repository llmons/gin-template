//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/llmons/gin-template/internal/base/conf"
	"github.com/llmons/gin-template/internal/base/data"
	"github.com/llmons/gin-template/internal/base/server"
)

func initApplication(
	debug bool,
	dbConf *data.Database,
	serverConf *conf.Server,
) (*gin.Engine, func(), error) {
	panic(wire.Build(
		server.ProviderSetServer,
	))
}
