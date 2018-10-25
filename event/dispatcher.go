package event

type Dispatcher interface {
	ListenFunc(event Event, fn interface{})
	Listen(event Event, listener Listener)
	Dispatch(event Event) (err error)
}
