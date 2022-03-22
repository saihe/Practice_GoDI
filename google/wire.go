//go:build wireinject
// +build wireinject

package google

import (
	"practice_godi/batch"
	"practice_godi/service"

	"github.com/google/wire"
)

func InitializeTask() batch.Task {
	wire.Build(
		batch.NewTask,
		service.NewService,
	)
	return batch.Task{}
}
