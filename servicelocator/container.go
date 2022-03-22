package servicelocator

import (
	"practice_godi/batch"
	"practice_godi/service"
)

var Container map[string]interface{}

func init() {
	Container = map[string]interface{}{
		"task": batch.NewTask(service.NewService()),
	}
}
