package di

import (
	"portfolio-backend/infra/repository"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRSSFeedRepository,
)
