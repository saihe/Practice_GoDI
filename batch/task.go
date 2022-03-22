package batch

import "practice_godi/service"

type Task struct {
	service service.SampleService
}

func NewTask(s service.SampleService) Task {
	return Task{service: s}
}

func (b Task) Execute(msg string) {
	b.service.Hello(msg)
}
