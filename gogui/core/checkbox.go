package core

type CheckBox struct {
	Wnd
	color     *Color
	pressFlag bool
	Checkable bool
	Checked   func(bool)
}

func (b *CheckBox) Draw() {
	if b.color == nil {
		b.color = &Color{}
	}
	b.draw(b.color)
}

func (b *CheckBox) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *CheckBox) draw(color *Color) {
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{255, 255, 255})
	if b.Hidden {
		return
	}
	b.Phy.DrawRect(b.X, b.Y+5, b.X+20, 25+b.Y, *color)
	if b.Checkable {
		// b.Phy.FillRect(b.X, b.Y+5, b.X+20, 25+b.Y, *color)
		b.Phy.DrawLine(b.X, b.Y+15, b.X+10, 20+b.Y, *color)
		b.Phy.DrawLine(b.X+10, 20+b.Y, b.X+20, b.Y, *color)
	} else {
		// b.Phy.DrawRect(b.X, b.Y+5, b.X+20, 25+b.Y, *color)
	}

	num := int(b.Width / FONT_WIDTH)
	n := len(b.Text)
	_text := b.Text
	if n > num {
		_text = b.Text[n-num : n]
	}
	b.Phy.DrawText(b.X+20+(b.Width-FONT_WIDTH*len(_text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, _text, *b.color)
}

func (b *CheckBox) SetColor(c *Color) {
	b.color = c
}

func (b *CheckBox) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *CheckBox) OnMouseDown(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.pressFlag = true
	}
	return false
}

func (b *CheckBox) OnMouseUp(x, y int) bool {
	b.pressFlag = false
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		root := b.findRoot()
		for _, child := range root.GetChildren() {
			child.SetFocus(false)
		}
		b.Checkable = !b.Checkable
		b.draw(b.color)
		if b.Checked != nil {
			b.Checked(b.Checkable)
		}
		return true
	}
	return false
}

func (b *CheckBox) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}
