package core

import (
	"reflect"
	"strings"
)

const (
	ATTR_VISIBLE  = 0x40000000
	ATTR_FOCUS    = 0x20000000
	ATTR_PRIORITY = 0x10000000 // Handle touch action at high priority
)

const (
	STATUS_NORMAL   = 0
	STATUS_PUSHED   = 1
	STATUS_FOCUSED  = 2
	STATUS_DISABLED = 3
)

const (
	MG_MOUSE_MOVE = 0
	MG_MOUSE_DOWN = 1
	MG_MOUSE_UP   = 2
)

const (
	TOUCH_DOWN = 1
	TOUCH_UP   = 2
)

type MouseEventFunc func(t, x, y int)
type KeyEventFunc func(key int)

type WndInterface interface {
	SetDisplay(phy *Display)
	SetPos(x, y, width, height int)
	SetFocus(f bool)
	Update()
	Draw()
	SetParent(p WndInterface)
	GetParent() WndInterface
	GetChildren() []WndInterface
	RemoveChild(child WndInterface)
	GetPriority() int
	Clear()
	findRoot() WndInterface
	OnMouseDownEvent(x, y int)
	OnMouseDown(x, y int) bool
	OnMouseUpEvent(x, y int)
	OnMouseUp(x, y int) bool
	OnMouseMoveEvent(x, y int)
	OnMouseMove(x, y int) bool
	OnKeyEvent(key int)
	OnKey(key int) bool
}

type Wnd struct {
	X       int
	Y       int
	Width   int
	Height  int
	X0      int
	Y0      int
	Width0  int
	Height0 int
	Phy     *Display
	Display
	Children       []WndInterface
	Text           string
	Hidden         bool
	Focus          bool
	Parent         WndInterface
	Priority       int
	DrawEvent      func()
	CustomKeyEvent func(int)
}

func NewWnd(window *Window, x, y, w, h int) *Wnd {
	phy := Display{W: window}
	wnd := &Wnd{Phy: &phy, X: x, Y: y, Width: w, Height: h}
	wnd.Display = phy
	wnd.Children = make([]WndInterface, 0)
	window.MouseEvent = wnd.MouseEvent
	window.KeyEvent = wnd.KeyEvent
	wnd.draw()
	return wnd
}

func (w *Wnd) Clear() {
	wi, hi := w.Phy.GetSize()
	w.Phy.FillRect(1, 1, wi, hi, Color{255, 255, 255})
}

func (w *Wnd) Update() {
	//w.Phy.FillRect(w.X+1, w.Y+1, w.Width-2, w.Height-1, Color{255, 255, 255})
	for _, child := range w.Children {
		if child.GetPriority() == 1 {
			child.Draw()
			child.Update()
		}
	}
	for _, child := range w.Children {
		if child.GetPriority() == 0 {
			child.Draw()
			child.Update()
		}
	}
}

func (w *Wnd) draw() {
	w.Phy.DrawRect(w.X+1, w.Y+1, w.Width-2, w.Height-1, Color{200, 200, 200})
}

func (w *Wnd) Draw() {
	if w.DrawEvent != nil {
		w.DrawEvent()
	}
}

func (w *Wnd) SetPos(x, y, width, height int) {
	w.Clear()
	w.X0 = w.X
	w.Y0 = w.Y
	w.Height0 = w.Height
	w.Width0 = w.Width

	w.X = x
	w.Y = y
	w.Height = height
	w.Width = width
}

func (w *Wnd) AddChild(child WndInterface) {
	w.Children = append(w.Children, child)
	child.SetDisplay(w.Phy)
	child.SetParent(w)
}

func (w *Wnd) RemoveChild(child WndInterface) {
	children := make([]WndInterface, 0)
	for _, item := range w.Children {
		name := reflect.TypeOf(item).String()
		flag := strings.Contains(name, "Messagebox")
		if !flag {
			children = append(children, item)
		} else {
			item = nil
		}
	}
	w.Children = children
}

func (w *Wnd) SetParent(p WndInterface) {
	w.Parent = p
}

func (w *Wnd) SetDisplay(phy *Display) {
	w.Phy = phy
	w.Display = *phy
}

func (w *Wnd) MouseEvent(t, x, y int) {
	switch t {
	case MG_MOUSE_MOVE:
		w.OnMouseMove(x, y)
		w.OnMouseMoveEvent(x, y)
		break
	case MG_MOUSE_DOWN:
		w.OnMouseDown(x, y)
		w.OnMouseDownEvent(x, y)
		break
	case MG_MOUSE_UP:
		w.OnMouseUp(x, y)
		w.OnMouseUpEvent(x, y)
	}
}

func (w *Wnd) KeyEvent(key int) {
	w.OnKey(key)
	w.OnKeyEvent(key)
}

func (w *Wnd) OnKeyEvent(key int) {
	for _, child := range w.Children {
		child.OnKey(key)
		child.OnKeyEvent(key)
	}
}

func (w *Wnd) OnKey(key int) bool {
	if w.CustomKeyEvent != nil {
		w.CustomKeyEvent(key)
	}
	return false
}

func (w *Wnd) OnMouseDownEvent(x, y int) {
	for _, child := range w.Children {
		child.OnMouseDown(x, y)
		child.OnMouseDownEvent(x, y)
	}
}

func (w *Wnd) OnMouseDown(x, y int) bool {
	return false
}

func (w *Wnd) OnMouseUpEvent(x, y int) {
	for priority := 1; priority >= 0; priority -= 1 {
		for _, child := range w.Children {
			if child.GetPriority() != priority {
				continue
			}
			if child.OnMouseUp(x, y) {
				return
			}
			child.OnMouseUpEvent(x, y)
		}
	}
}

func (w *Wnd) OnMouseUp(x, y int) bool {
	return false
}

func (w *Wnd) OnMouseMoveEvent(x, y int) {
	for _, child := range w.Children {
		child.OnMouseMove(x, y)
		child.OnMouseMoveEvent(x, y)
	}
}

func (w *Wnd) OnMouseMove(x, y int) bool {

	return false
}

func (w *Wnd) SetFocus(f bool) {
	w.Focus = f
	w.Draw()
}

func (w *Wnd) GetParent() WndInterface {
	return w.Parent
}

func (w *Wnd) GetChildren() []WndInterface {
	return w.Children
}

func (w *Wnd) findRoot() WndInterface {
	var s WndInterface = w
	for {
		p := s.GetParent()
		if p.GetParent() == nil {
			return p
		}
		s = p
	}
}

func (w *Wnd) GetPriority() int {
	return w.Priority
}
