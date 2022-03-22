package main

import (
	"practice_godi/batch"
	"practice_godi/google"
	"practice_godi/manual"
	"practice_godi/service"
	"practice_godi/servicelocator"

	"go.uber.org/dig"
)

func main() {
	// Without container.
	b1 := batch.NewTask(service.NewService())
	b1.Execute("コンテナなし")

	// With manual container.
	manualContainer := manual.NewContainer()
	manualContainer.Task.Execute("手動コンテナ使用")

	// With google/wire
	b2 := google.InitializeTask()
	b2.Execute("Google Wire 使用")

	// With Rakusul dig
	digContainer := dig.New()
	digContainer.Provide(service.NewService)
	digContainer.Provide(batch.NewTask)
	digContainer.Invoke(func(task batch.Task) {
		task.Execute("Raksul dig 使用")
	})

	// With Serivce Locator
	b3 := servicelocator.Container["task"].(batch.Task)
	b3.Execute("サービスロケーターパターンで実装")
}
