package singleton

import "sync"

var (
	lockInstance *guidGenerator
	mutex        = &sync.Mutex{}
)

func GetLockInstance() *guidGenerator {
	if lockInstance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if lockInstance == nil {
			lockInstance = &guidGenerator{}
		}
	}
	return lockInstance
}
