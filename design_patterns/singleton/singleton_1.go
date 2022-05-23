package singleton

// 懒汉模式singleton，非线程安全
type singleTon1 struct{

}

var instance *singleTon1

func getInstance1()*singleTon1 {
	if instance == nil{
		instance = &singleTon1{}
	}
	return instance
}