package core

func inRect(x0, y0, x, y, w, h int) bool {
	if x0 >= x && x0 <= x+w && y0 >= y && y0 <= y+h {
		return true
	}
	return false
}
