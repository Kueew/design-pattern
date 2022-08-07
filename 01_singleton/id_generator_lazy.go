package singleton

import "sync"

var (
	layzInstance *guidGenerator
	once         = &sync.Once{}
)

func GetLazyInstance() *guidGenerator {
	if layzInstance == nil {
		once.Do(func() {
			layzInstance = &guidGenerator{}
		})
	}
	return layzInstance
}
