package service

import "fmt"

type SampleService interface {
	Hello(msg string)
}

type ImplementsService struct{}

func NewService() SampleService {
	return &ImplementsService{}
}

func (*ImplementsService) Hello(msg string) {
	fmt.Printf("Hello %s.\n", msg)
}
