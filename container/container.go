package container

//	Go 没有泛型，所有接口的方法参数都不指定类型
//	在特定场景下的容器时也该如此
type Container interface {

	//	Register a binding to the container
	//	注册一个 Binding 对象到容器里
	BindAny(name string, any interface{})

	//	Register an existing instance as shared in the container
	//	注册一个 实例对象 到容器，该实例会作为一个 shared binding 存在
	BindInstance(string, interface{})

	//	Register a shared binding to the container
	Singleton(string, interface{})

	//
	Alias(string, string)

	//	从已注册的 Binding 里生成实例
	Make(string, ...interface{}) interface{}

	//	判断是否被绑定过
	IsBound(string) bool

	//	判断是否被解析过
	IsResolved(string) bool

	//	判断是否共享
	IsShared(string) bool

	//	判断是否别名
	IsAlias(string) bool

	//	注册解析事件
	BeforeResolving(string, interface{})

	//	注册解析后的事件
	AfterResolved(string, interface{})
}
