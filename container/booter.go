package container

//	首次实例化前调用一次
type Booter interface {
	Boot()
}
