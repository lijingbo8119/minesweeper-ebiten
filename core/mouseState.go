package core

type mouseHandler func(c Coordinate)

type mouseHandlers []mouseHandler

func (this mouseHandlers) each(closure func(h mouseHandler)) {
	for _, h := range this {
		closure(h)
	}
}

func (this mouseHandlers) do(c Coordinate) {
	this.each(func(h mouseHandler) { h(c) })
}

type MouseState struct {
	leftMouseDownCoordinate  Coordinate
	leftMouseUpCoordinate    Coordinate
	rightMouseDownCoordinate Coordinate
	rightMouseUpCoordinate   Coordinate

	leftMouseDownHandler      mouseHandlers
	leftMouseUpHandler        mouseHandlers
	rightMouseDownHandler     mouseHandlers
	rightMouseUpHandler       mouseHandlers
	leftRightMouseDownHandler mouseHandlers
	leftRightMouseUpHandler   mouseHandlers
	resetHandler              mouseHandlers
}

func (this *MouseState) ResetAllCoordinate(c Coordinate) *MouseState {
	this.leftMouseDownCoordinate = NewEmptyCoordinate()
	this.leftMouseUpCoordinate = NewEmptyCoordinate()
	this.rightMouseDownCoordinate = NewEmptyCoordinate()
	this.rightMouseUpCoordinate = NewEmptyCoordinate()
	this.resetHandler.do(c)
	return this
}

func (this *MouseState) LeftMouseDown(c Coordinate) *MouseState {
	this.leftMouseDownCoordinate = c
	if this.leftMouseDownCoordinate.Equal(this.rightMouseDownCoordinate) {
		this.leftRightMouseDownHandler.do(c)
		return this
	}
	if !this.rightMouseDownCoordinate.IsEmpty() {
		this.rightMouseDownCoordinate = NewEmptyCoordinate()
	}
	this.leftMouseDownHandler.do(c)
	return this
}

func (this *MouseState) LeftMouseUp(c Coordinate) *MouseState {
	this.leftMouseUpCoordinate = c
	if !this.leftMouseUpCoordinate.Equal(this.leftMouseDownCoordinate) {
		this.ResetAllCoordinate(c)
		return this
	}
	if this.leftMouseUpCoordinate.Equal(this.rightMouseUpCoordinate) {
		this.leftRightMouseUpHandler.do(c)
		this.ResetAllCoordinate(c)
		return this
	}
	if this.rightMouseDownCoordinate.IsEmpty() {
		this.leftMouseUpHandler.do(c)
		this.ResetAllCoordinate(c)
		return this
	}
	if this.leftMouseUpCoordinate.Equal(this.leftMouseDownCoordinate) {
		return this
	}
	this.ResetAllCoordinate(c)
	return this
}

func (this *MouseState) RightMouseDown(c Coordinate) *MouseState {
	this.rightMouseDownCoordinate = c
	if this.rightMouseDownCoordinate.Equal(this.leftMouseDownCoordinate) {
		this.leftRightMouseDownHandler.do(c)
		return this
	}
	if !this.leftMouseDownCoordinate.IsEmpty() {
		this.leftMouseDownCoordinate = NewEmptyCoordinate()
	}
	this.rightMouseDownHandler.do(c)
	return this
}

func (this *MouseState) RightMouseUp(c Coordinate) *MouseState {
	this.rightMouseUpCoordinate = c
	if !this.rightMouseUpCoordinate.Equal(this.rightMouseDownCoordinate) {
		this.ResetAllCoordinate(c)
		return this
	}
	if this.rightMouseUpCoordinate.Equal(this.leftMouseUpCoordinate) {
		this.leftRightMouseUpHandler.do(c)
		this.ResetAllCoordinate(c)
		return this
	}
	if this.leftMouseDownCoordinate.IsEmpty() {
		this.rightMouseUpHandler.do(c)
		this.ResetAllCoordinate(c)
		return this
	}
	if this.rightMouseUpCoordinate.Equal(this.leftMouseDownCoordinate) {
		return this
	}
	this.ResetAllCoordinate(c)
	return this
}

func (this *MouseState) RegisterLeftMouseDownHandler(h ...mouseHandler) *MouseState {
	this.leftMouseDownHandler = h
	return this
}

func (this *MouseState) RegisterLeftMouseUpHandler(h ...mouseHandler) *MouseState {
	this.leftMouseUpHandler = h
	return this
}

func (this *MouseState) RegisterRightMouseDownHandler(h ...mouseHandler) *MouseState {
	this.rightMouseDownHandler = h
	return this
}

func (this *MouseState) RegisterRightMouseUpHandler(h ...mouseHandler) *MouseState {
	this.rightMouseUpHandler = h
	return this
}

func (this *MouseState) RegisterLeftRightMouseDownHandler(h ...mouseHandler) *MouseState {
	this.leftRightMouseDownHandler = h
	return this
}

func (this *MouseState) RegisterLeftRightMouseUpHandler(h ...mouseHandler) *MouseState {
	this.leftRightMouseUpHandler = h
	return this
}

func (this *MouseState) RegisterResetHandler(h ...mouseHandler) *MouseState {
	this.resetHandler = h
	return this
}

func NewMouseState() *MouseState {
	return new(MouseState).ResetAllCoordinate(NewEmptyCoordinate())
}
