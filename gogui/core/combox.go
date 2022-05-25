package core

type Combox struct {
	Wnd
	color       *Color
	pressFlag   bool
	normalColor *Color
	pressColor  *Color
	items       []string
	orgBytes    []byte
	boxFlag     bool
	itemHeight  int
	Click       func(string)
}

func (b *Combox) Draw() {
	if b.color == nil {
		b.itemHeight = 30
		b.Priority = 1
		b.color = &Color{}
		b.pressColor = &Color{R: 0, G: 255, B: 0}
		b.normalColor = &Color{}
	}
	if b.Focus {
		b.draw(b.pressColor)
	} else {
		b.draw(b.normalColor)
		num := len(b.items)
		if b.orgBytes != nil {
			b.Phy.SetPixels(b.X, b.Height+b.Y+num*b.itemHeight+1, b.Width, (num+1)*b.itemHeight, b.orgBytes)

		}
	}
}

func (b *Combox) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *Combox) drawTriangle(color *Color) {
	b.Phy.DrawLine(b.Width+b.X-13, b.Height+b.Y-20, b.Width+b.X-8, b.Height+b.Y-10, *color)
	b.Phy.DrawLine(b.Width+b.X-3, b.Height+b.Y-20, b.Width+b.X-8, b.Height+b.Y-10, *color)
	b.Phy.DrawHLine(b.Width+b.X-13, b.Width+b.X-2, b.Height+b.Y-20, *color)
}

func (b *Combox) drawList() {
	if !b.Focus {
		return
	}
	b.boxFlag = true
	// b.items = []string{"item1", "item2", "item3"}
	num := len(b.items)
	if num < 1 {
		return
	}
	// b.orgBytes = b.Phy.GetPixels(b.X, b.Height+b.Y+num*itemHeight, b.Width, (num+1)*itemHeight)
	for i, item := range b.items {
		b.Phy.FillRect(b.X, b.Height+b.Y+i*b.itemHeight+1, b.X+b.Width, b.Height+b.Y+(i+1)*b.itemHeight, Color{R: 100, G: 200})
		b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(item))/2, b.Height+b.Y+i*b.itemHeight+(b.itemHeight-FONT_HEIGHT)/2, item, *b.color)
	}
}

func (b *Combox) draw(color *Color) {
	if b.boxFlag {
		b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y+(len(b.items)+1)*b.itemHeight, Color{255, 255, 255})
		b.boxFlag = false
	}
	if b.Hidden {
		return
	}
	b.drawList()
	b.drawTriangle(color)
	b.Phy.DrawRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, *b.normalColor)
	num := int(b.Width / FONT_WIDTH)
	n := len(b.Text)
	_text := b.Text
	if n > num {
		_text = b.Text[n-num : n]
	}
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(_text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, _text, *b.color)
}

func (b *Combox) SetColor(c *Color) {
	b.color = c
}

func (b *Combox) SetText(text string) {
	if b.items == nil {
		b.items = make([]string, 0)
	}
	b.items = append(b.items, text)
	b.Text = text
	b.Draw()
}

func (b *Combox) OnMouseDown(x, y int) bool {
	if inRect(x, y, b.X, b.Y, b.Width, b.Height) {
		b.pressFlag = true
	}
	return false
}

func (b *Combox) click(y int) {
	y = y - b.Y - b.Height
	if y <= 0 || len(b.items) < 1 {
		return
	}
	b.Focus = false
	index := int(y / b.itemHeight)
	if index < len(b.items) {
		b.Text = b.items[index]
		if b.Click != nil {
			b.Click(b.items[index])
		}
	}
}

func (b *Combox) AddItem(text string) {
	b.items = append(b.items, text)
}

func (b *Combox) OnMouseUp(x, y int) bool {

	h := 0
	if b.Focus {
		h = b.Height + len(b.items)*b.itemHeight
	} else {
		h = b.Height
	}
	if inRect(x, y, b.X, b.Y, b.Width, h) {
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

		b.SetFocus(b.Focus)
		b.draw(b.pressColor)
		b.pressFlag = false
		b.click(y)
		if !b.Focus {
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
		}
		return true
	}
	b.pressFlag = false
	return false
}

func (b *Combox) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}
