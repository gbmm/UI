package core

type GridLayout struct {
	Parent    *Wnd
	Rows      int
	Columns   int
	Hinterval int
	Vinterval int
}

func (b *GridLayout) AddChild(obj WndInterface, row, col, rowSpan, colSpan int) {
	x0 := 1
	y0 := 1
	w := (b.Parent.Width - x0 - 20) / b.Columns
	h := (b.Parent.Height - y0 - 45) / b.Rows
	obj.SetPos(x0+col*w, y0+row*h, colSpan*w-b.Hinterval, rowSpan*h-b.Vinterval)
	b.Parent.AddChild(obj)
	b.Parent.Update()
}

func (b *GridLayout) SetInterval(v, h int) {
	b.Vinterval = v
	b.Hinterval = h
}

func NewGridLayout(p *Wnd, row, col int) *GridLayout {
	layout := &GridLayout{Parent: p, Rows: row, Columns: col}
	layout.SetInterval(2, 2)
	return layout
}
