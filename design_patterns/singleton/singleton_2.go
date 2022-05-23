package singleton

import "sync"

// 带锁的单例，线程安全
type singleTon2 struct {

}

var instance2 *singleTon2
var mu sync.Mutex

func getInstance2()*singleTon2{
	mu.Lock()
	defer mu.Unlock()

	if instance2 == nil{
		instance2 = &singleTon2{}
	}
	return instance2
}

