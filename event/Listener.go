package event

type Listener interface {
	Handle(Event) error
}
