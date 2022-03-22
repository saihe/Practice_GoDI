package manual

import (
	"practice_godi/batch"
	"practice_godi/service"
)

type Container struct {
	Task batch.Task
}

func NewContainer() Container {
	c := Container{}
	c.Initialize()
	return c
}

func (c *Container) Initialize() {
	c.Task = batch.NewTask(
		service.NewService(),
	)
}
