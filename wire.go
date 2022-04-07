//go:build wireinject
// +build wireinject

package main

import (
	"rest-api/api"
	"rest-api/repo"
	"rest-api/service"

	"github.com/google/wire"
)

func initApi() api.HiAPI {
	wire.Build(api.HiAPISet, service.HiServiceSet, repo.MessageRepoSet)
	return api.HiAPI{}
}
