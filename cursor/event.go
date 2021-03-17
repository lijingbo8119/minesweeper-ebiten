package cursor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"sync"
)

type Event func(s *Status) bool

type events []Event

func (this events) each(f func(e Event)) {
	for _, event := range this {
		f(event)
	}
}

var eventMu sync.RWMutex

var eventsInstance = map[ebiten.MouseButton]map[Action]events{}

func RegisterEvent(b ebiten.MouseButton, a Action, f Event) {
	eventMu.Lock()
	defer eventMu.Unlock()

	if _, ok := eventsInstance[b]; !ok {
		eventsInstance[b] = map[Action]events{}
	}

	if _, ok := eventsInstance[b][a]; !ok {
		eventsInstance[b][a] = events{}
	}

	eventsInstance[b][a] = append(eventsInstance[b][a], f)
}
