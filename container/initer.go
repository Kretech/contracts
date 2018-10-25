package container

//	每次实例化后都会被调用
type Initer interface {
	Init()
}
