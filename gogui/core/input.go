package core

type Input struct {
	Wnd
	color       *Color
	pressFlag   bool
	normalColor *Color
	pressColor  *Color
}

func (b *Input) Draw() {
	if b.color == nil {
		b.color = &Color{}
		b.pressColor = &Color{R: 0, G: 255, B: 0}
		b.normalColor = &Color{}
	}
	if b.Focus {
		b.draw(b.pressColor)
	} else {
		b.draw(b.normalColor)
	}
}

func (b *Input) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *Input) draw(color *Color) {
	//if b.X0 != b.X || b.Y0 != b.Y || b.Height0 != b.Height || b.Width0 != b.Width {
	//	b.Phy.FillRect(b.X0, b.Y0, b.Width0+b.X0+1, b.Height0+b.Y0+1, Color{255, 255, 255})
	//}
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{255, 255, 255})
	if b.Hidden {
		return
	}
	b.Phy.DrawRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, *color)
	num := int(b.Width / FONT_WIDTH)
	n := len(b.Text)
	_text := b.Text
	if n > num {
		_text = b.Text[n-num : n]
	}
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(_text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, _text, *b.color)
}

func (b *Input) SetColor(c *Color) {
	b.color = c
}

func (b *Input) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *Input) OnMouseDown(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.pressFlag = true
	}
	return false
}

func (b *Input) OnMouseUp(x, y int) bool {
	b.pressFlag = false
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		root := b.findRoot()
		for _, child := range root.GetChildren() {
			child.SetFocus(false)
		}
		b.Focus = true
		b.SetFocus(true)
		b.draw(b.pressColor)
		return true
	}
	return false
}

func (b *Input) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}

func (b *Input) OnKey(key int) bool {
	if !b.Focus {
		return false
	}
	if key == 8 && len(b.Text) > 0 {
		b.Text = b.Text[0 : len(b.Text)-1]
		b.Draw()
	} else if key > 47 {
		b.Text += string(rune(key))
		b.Draw()
	}
	return true
}
