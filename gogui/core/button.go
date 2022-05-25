package core

type Button struct {
	Wnd
	normalColor *Color
	pressColor  *Color
	Click       func()
	pressFlag   bool
}

func (b *Button) Draw() {
	if b.normalColor == nil {
		b.normalColor = &Color{B: 150, G: 150, R: 150}
		b.pressColor = &Color{B: 0, G: 150, R: 150}
	}
	b.draw(b.normalColor)
}

func (b *Button) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *Button) draw(color *Color) {
	if b.Hidden {
		b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{255, 255, 255})
		return
	}
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, *color)
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(b.Text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, b.Text, Color{})

}

func (b *Button) SetColor(normalColor, pressColor *Color) {
	b.normalColor = normalColor
	b.pressColor = pressColor
}

func (b *Button) OnMouseDown(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.draw(b.pressColor)
		b.pressFlag = true
	}
	return false
}

func (b *Button) OnMouseUp(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.draw(b.normalColor)
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
		b.Focus = true
		if b.pressFlag && b.Click != nil {
			b.Click()
		}
		b.pressFlag = false
		return true
	}
	b.pressFlag = false
	return false
}

func (b *Button) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *Button) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}
