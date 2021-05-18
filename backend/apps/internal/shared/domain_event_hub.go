package shared

import "log"

type DomainEventHandlerFunc func(DomainEvent)

type DomainEventHub struct {
	eventHandlers map[string][]DomainEventHandlerFunc
	events        chan DomainEvent
}

func NewDomainEventHub() *DomainEventHub {
	return &DomainEventHub{
		eventHandlers: make(map[string][]DomainEventHandlerFunc),
		events:        make(chan DomainEvent),
	}
}

func (hub *DomainEventHub) Raise(events ...DomainEvent) {
	go func() {
		for _, event := range events {
			hub.events <- event
		}
	}()
}

func (hub *DomainEventHub) Register(event DomainEvent, handler DomainEventHandlerFunc) {
	if _, ok := hub.eventHandlers[event.Name()]; !ok {
		hub.eventHandlers[event.Name()] = make([]DomainEventHandlerFunc, 0)
	}
	hub.eventHandlers[event.Name()] = append(hub.eventHandlers[event.Name()], handler)
}

func (hub *DomainEventHub) Start() {
	go func() {
		for {
			event := <-hub.events
			log.Printf("PROCESSING EVENT %s", event.Name())
			for _, handler := range hub.eventHandlers[event.Name()] {
				go handler(event)
			}
		}
	}()
}
