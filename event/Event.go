package event

type Event interface{}

type EventKey = string

type EventKeyHolder interface {
	EventKey() EventKey
}
