package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

func (e *EventDispatcher) Clear() error {
	e.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (e *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (e *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for i, h := range e.handlers[eventName] {
			if h == handler {
				// remove handler from slice
				// [:i] -> from beginning to i
				// [i+1:] -> from i+1 to end
				// append -> concatenate slices and return new slice
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
