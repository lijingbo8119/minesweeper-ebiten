package cursor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action int

const (
	ActionRelease Action = iota
	ActionPress
)

func GetAction(b ebiten.MouseButton) Action {
	var (
		valInstance *Val
		val         *Status
	)

	switch b {
	case ebiten.MouseButtonLeft:
		valInstance = leftButtonValInstance
	case ebiten.MouseButtonRight:
		valInstance = rightButtonValInstance
	}

	val = valInstance.Get().(*Status)

	if ebiten.IsMouseButtonPressed(b) && val.Action == ActionRelease {
		val.Position = GetPosition()
		val.Action = ActionPress
		valInstance.Set(val)
		eventsInstance[b][val.Action].each(func(e Event) { e(val) })
		return ActionPress
	}

	if ebiten.IsMouseButtonPressed(b) && val.Action == ActionPress {
		val.Position = GetPosition()
		valInstance.Set(val)
		return ActionPress
	}

	if inpututil.IsMouseButtonJustReleased(b) && val.Action == ActionPress {
		val.Position = GetPosition()
		val.Action = ActionRelease
		valInstance.Set(val)
		eventsInstance[b][val.Action].each(func(e Event) { e(val) })
		return ActionRelease
	}

	val.Position = GetPosition()
	valInstance.Set(val)
	return ActionRelease
}
