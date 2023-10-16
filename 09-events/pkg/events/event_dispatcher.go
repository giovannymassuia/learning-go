package events

import (
	"errors"
	"sync"
)

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
		wg := &sync.WaitGroup{} // creates wait group to wait for all handlers to finish
		for _, handler := range handlers {
			wg.Add(1) // add 1 to wait group

			go handler.Handle(event, wg)

			//go func() {
			//	handler.Handle(event)
			//	wg.Done() // remove 1 from wait group
			//}()
		}
		wg.Wait() // wait for all handlers to finish
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
