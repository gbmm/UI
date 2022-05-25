package core

type Messagebox struct {
	Wnd
	color     *Color
	Title     string
	pressFlag bool
}

func (b *Messagebox) Draw() {
	if b.color == nil {
		b.color = &Color{}
		b.Priority = 1
	}
	b.draw(b.color)
}

func (b *Messagebox) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *Messagebox) draw(color *Color) {
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X+2, b.Height+b.Y, Color{255, 255, 255})
	if b.Hidden {
		return
	}
	b.Phy.DrawRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{200, 200, 200})
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Y+30, Color{200, 200, 200})
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(b.Title))/2, b.Y+10, b.Title, *b.color)
	// b.Phy.DrawHLine(b.X, b.X+b.Width, b.Y+30, *b.color)
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(b.Text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, b.Text, *b.color)
}

func (b *Messagebox) SetColor(c *Color) {
	b.color = c
}

func (b *Messagebox) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *Messagebox) OnMouseDown(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.pressFlag = true
	}
	return false
}

func (b *Messagebox) OnMouseUp(x, y int) bool {
	b.pressFlag = false
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.Hidden = true
		b.draw(b.color)
		root := b.findRoot()
		for _, child := range root.GetChildren() {
			if child.GetPriority() == 1 {
				child.SetFocus(false)
			}
		}
		for _, child := range root.GetChildren() {
			if child.GetPriority() == 0 {
				child.SetFocus(false)
			}
		}
		root.RemoveChild(b)
		return true
	}
	return false
}

func MessageBoxInfo(w *Wnd, title, info string) {
	width := len(info)*20 + 30
	height := 60 + 40
	x := (w.Width - width) / 2
	y := (w.Height - height) / 4
	box := &Messagebox{Wnd: Wnd{X: x, Y: y, Width: width, Height: height, Text: info}, Title: title}
	w.AddChild(box)
	w.Update()
}
