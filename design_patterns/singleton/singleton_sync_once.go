package singleton

import "sync"

// 采用golang内部的sync.Once实现的单例模式
type singleTonSync struct{

}

var instanceSync *singleTonSync
var once sync.Once

func getInstanceSync() *singleTonSync{
	var initialize func()
	initialize = func(){
		if instanceSync == nil{
			instanceSync = &singleTonSync{}
		}
	}

	once.Do(initialize)
	return instanceSync
}