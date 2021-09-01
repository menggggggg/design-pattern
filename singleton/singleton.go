package singleton

import "sync"

var (
	instance *singleObject
	once     sync.Once
)

type singleObject struct {
}

func (s *singleObject) message() string {
	return "singleton object!"
}

func GetSingleInstance() *singleObject {
	once.Do(func() {
		instance = &singleObject{}
	})
	return instance
}
