package singleton

import "sync"

var (
	singleton *Singleton
	once      sync.Once
)

func init() {
	//	singleton = &Singleton{} 饿汉模式
}

type Singleton struct {
}

func GetSingleInstance() *Singleton {
	// 懒汉模式
	if singleton == nil {
		once.Do(func() {
			singleton = &Singleton{}
		})
	}
	return singleton
}
