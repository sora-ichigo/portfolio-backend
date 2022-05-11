//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func NewApp() (*App, error) {
	wire.Build(wire.Struct(new(App), "*"), RepositorySet, HandlerSet)

	return nil, nil
}
