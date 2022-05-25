package core

type Label struct {
	Wnd
	color      *Color
	imageData  *[]byte
	imageWidth int
}

func (b *Label) Draw() {
	if b.color == nil {
		b.color = &Color{}
	}
	b.draw(b.color)
}

func (b *Label) SetVisible(flag bool) {
	b.Hidden = !flag
	b.Draw()
}

func (b *Label) draw(color *Color) {
	b.Phy.FillRect(b.X, b.Y, b.Width+b.X, b.Height+b.Y, Color{255, 255, 255})
	if b.Hidden {
		return
	}
	if b.imageData != nil {
		b.Phy.SetImagePixels(b.X, b.Y, b.Width, b.Height, *b.imageData, b.imageWidth)
	}
	if b.DrawEvent != nil {
		b.DrawEvent()
	}
	b.Phy.DrawText(b.X+(b.Width-FONT_WIDTH*len(b.Text))/2, b.Y+(b.Height-FONT_HEIGHT)/2, b.Text, *b.color)
}

func (b *Label) SetColor(c *Color) {
	b.color = c
}

func (b *Label) SetText(text string) {
	b.Text = text
	b.Draw()
}

func (b *Label) SetFocus(f bool) {
	b.Focus = f
	b.Draw()
}

func (b *Label) SetImage(bmpPath string) {
	info, data := ReaaBmp(bmpPath)
	w := int(info.BiWidth)
	h := int(info.BiHeight)
	if int(info.BiWidth) > b.Width {
		w = b.Width
	}
	if int(info.BiHeight) > b.Height {
		h = b.Height
	}
	b.Phy.SetImagePixels(b.X, b.Y, w, h, *data, int(info.BiWidth))
	b.imageData = data
	b.imageWidth = int(info.BiWidth)
}
