package core


type Color struct {
	R uint8
	G uint8
	B uint8
}

type Display struct {
	W          *Window
	showIndex  int
	writeIndex int
}

func (display *Display) GetSize() (int, int) {
	if display ==nil || display.W == nil{
		return 0,0
	}
	return int(display.W.BmpInfo.BmiHeader.BiWidth), int(display.W.BmpInfo.BmiHeader.BiHeight)
}

func (display *Display) changeXY(x, y int) (int, int) {
	y = int(display.W.BmpInfo.BmiHeader.BiHeight) - y
	if y < 45 {
		y = 45
	}
	if x > (int(display.W.BmpInfo.BmiHeader.BiWidth) - 20) {
		x = int(display.W.BmpInfo.BmiHeader.BiWidth) - 20
	}
	return x, y
}

func (display *Display) GetPixels(x, y, w, h int) []byte {
	x, y = display.changeXY(x, y)
	line := int(display.W.Line)
	index := 0
	bs := make([]byte, (w+1)*4*h)
	for y0 := y; y0 < y+h; y0++ {
		for x0 := x; x0 < x+w; x0++ {
			t := y0*line + x0*3
			bs[index] = display.W.Bytes[t]
			bs[index+1] = display.W.Bytes[t+1]
			bs[index+2] = display.W.Bytes[t+2]
			index += 3
		}
	}
	return bs
}

func (display *Display) SetPixels(x, y, w, h int, bs []byte) {
	x, y = display.changeXY(x, y)
	line := int(display.W.Line)
	index := 0
	for y0 := y; y0 < y+h; y0++ {
		for x0 := x; x0 < x+w; x0++ {
			t := y0*line + x0*3
			display.W.Bytes[t] = bs[index]
			display.W.Bytes[t+1] = bs[index+1]
			display.W.Bytes[t+2] = bs[index+2]
			index += 3

		}
	}
}

func (display *Display) SetImagePixels(x, y, w, h int, bs []byte, imageWidth int) {
	x, y = display.changeXY(x, y)
	line := int(display.W.Line)
	line1 := (imageWidth*(24/8) + 3) / 4 * 4
	lines := len(bs) / line1
	if h > lines {
		h = lines
	}
	for y0 := y - h; y0 < y; y0++ {
		for x0 := x; x0 < x+w; x0++ {
			t := y0*line + x0*3
			t1 := (y0-y+h)*line1 + (x0-x)*3
			if (x0-x+1)*3 < line1 {
				display.W.Bytes[t] = bs[t1]
				display.W.Bytes[t+1] = bs[t1+1]
				display.W.Bytes[t+2] = bs[t1+2]
			}
		}
	}
}

func (display *Display) DrawPixel(y, line, x int, color Color) {
	x, y = display.changeXY(x, y)
	index := y*line + x*3
	display.W.Bytes[index] = color.B
	display.W.Bytes[index+1] = color.G
	display.W.Bytes[index+2] = color.R
	display.writeIndex = display.writeIndex + 1
	display.W.UpdateFlag = true
}

func (display *Display) DrawHLine(x1, x2, y int, color Color) {
	line := int(display.W.Line)
	for col := x1; col < x2; col = col + 1 {
		display.DrawPixel(y, line, col, color)
	}
}

func (display *Display) DrawVLine(y1, y2, x int, color Color) {
	line := int(display.W.Line)
	for y := y1; y < y2; y = y + 1 {
		display.DrawPixel(y, line, x, color)
	}
}

func (display *Display) DrawLine(x1, y1, x2, y2 int, color Color) {
	var dx, dy, x, y, e int
	line := int(display.W.Line)
	if x1 > x2 {
		dx = x1 - x2
	} else {
		dx = x2 - x1
	}
	if y1 > y2 {
		dy = y1 - y2
	} else {
		dy = y2 - y1
	}
	if (dx > dy && x1 > x2) || (dx <= dy && y1 > y2) {
		x = x2
		y = y2
		x2 = x1
		y2 = y1
		x1 = x
		y1 = y
	}

	x = x1
	y = y1
	if dx > dy {
		e = dy - dx/2
		for x1 <= x2 {
			x1 += 1
			e += dy
			display.DrawPixel(y1, line, x1, color)
			if e > 0 {
				e -= dx
				if y > y2 {
					y1 -= 1
				} else {
					y1 += 1
				}
			}
		}
	} else {
		e = dx - dy/2
		for y1 <= y2 {
			y1 += 1
			e += dx
			display.DrawPixel(y1, line, x1, color)
			if e > 0 {
				e -= dy
				if x > x2 {
					x1 -= 1
				} else {
					x1 += 1
				}
			}
		}
	}
}

func (display *Display) DrawRect(x1, y1, x2, y2 int, color Color) {
	display.DrawHLine(x1, x2, y1, color)
	display.DrawHLine(x1, x2, y2, color)
	display.DrawVLine(y1, y2, x1, color)
	display.DrawVLine(y1, y2, x2, color)
}

func (display *Display) FillRect(x1, y1, x2, y2 int, color Color) {
	// display.DrawRect(x1, y1, x2, y2, color)
	for h := y1; h <= y2; h = h + 1 {
		display.DrawHLine(x1, x2, h, color)
	}
}

func (display *Display) drawChar(x, y int, s string, color Color) {
	fontPix, ok := FONT[s]
	r := 255
	g := 255
	b := 255
	for y0 := 0; y0 < 25; y0++ {
		for x0 := 0; x0 < 20; x0++ {
			if !ok {
				r = 255
				g = 255
				b = 255
			} else {
				r = fontPix[x0+20*y0]
				g = fontPix[x0+20*y0]
				b = fontPix[x0+20*y0]
			}

			if r >= 0 && r <= 130 {
				r = int(color.R)
				b = int(color.B)
				g = int(color.G)
				display.DrawPixel(y0+y, int(display.W.Line), x0+x, Color{R: uint8(r), G: uint8(g), B: uint8(b)})
			}
		}
	}
}

func (display *Display) DrawText(x, y int, text string, color Color) {
	for i, s := range text {
		display.drawChar(x+i*FONT_WIDTH, y, string(s), color)
	}
}

func (display *Display) IsUpdate() bool {
	return display.writeIndex > 0
}

func (display *Display) DUpdate() {
	display.writeIndex = 0
}
