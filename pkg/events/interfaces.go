package events

import (
	"sync"
	"time"
)

type EventInterface interface {  // interface que sera implementada pelos eventos
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface { // manipulador dos eventos
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface { // disparador de eventos
	Register(eventName string, handler EventHandlerInterface) error  // registrar um evento
	Dispatch(event EventInterface) error  // disparar um evento
	Remove(eventName string, handler EventHandlerInterface) error // remover um evento
	Has(eventName string, handler EventHandlerInterface) bool // verificar se um evento existe
	Clear() error // limpar todos os eventos
}


