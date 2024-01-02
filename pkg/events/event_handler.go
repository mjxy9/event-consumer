package events

type EventHandler interface {
	Consume() error
}
