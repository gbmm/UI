package core

import (
	"fmt"
	"github.com/lxn/win"
	"syscall"
	"time"
	"unsafe"
)

type Window struct {
	Hwnd       win.HWND
	hdc        win.HDC
	BmpInfo    win.BITMAPINFO
	mMemdc     win.HDC
	Bytes      *[1 << 30]byte
	Hbitmap    win.HBITMAP
	Line       int32
	UpdateFlag bool
	MouseEvent MouseEventFunc
	KeyEvent   KeyEventFunc
	Width      int
	Height     int
}

var GHandWindow map[win.HWND]*Window

func CreateWindow(width, height int32, title string) *Window {
	if GHandWindow == nil {
		GHandWindow = make(map[win.HWND]*Window)
	}
	hInst := win.GetModuleHandle(nil)
	hIcon := win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	hCursor := win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))

	var wc win.WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = syscall.NewCallback(wndProc)
	wc.HInstance = hInst
	wc.HIcon = hIcon
	wc.HCursor = hCursor
	wc.HbrBackground = win.COLOR_WINDOW + 1
	wc.LpszClassName = syscall.StringToUTF16Ptr(title)
	wc.Style = win.CS_HREDRAW | win.CS_VREDRAW
	win.RegisterClassEx(&wc)

	hWnd := win.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr(title),
		syscall.StringToUTF16Ptr(title),
		win.WS_OVERLAPPEDWINDOW,
		400, 100, width, height, 0, 0, hInst, nil)

	win.ShowWindow(hWnd, win.SW_SHOW)
	win.UpdateWindow(hWnd)
	w := initWindow(hWnd, width, height)
	GHandWindow[hWnd] = w
	return w
}

func wndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {
	window := GHandWindow[hwnd]
	switch msg {
	case win.WM_SIZE:
		if w, ok := GHandWindow[hwnd]; ok {
			w.UpdateFlag = true
		}
		break
	case win.WM_CHAR:
		window.KeyEvent(int(wParam))
		break
	case win.WM_MOUSEMOVE:
		y := int(int32(lParam) >> 16)
		x := int(int32(lParam) & 0xFFFF)
		window.MouseEvent(MG_MOUSE_MOVE, x, y)
		break
	case win.WM_RBUTTONDOWN, win.WM_LBUTTONDOWN:
		y := int(int32(lParam) >> 16)
		x := int(int32(lParam) & 0xFFFF)
		window.MouseEvent(MG_MOUSE_DOWN, x, y)
		break
	case win.WM_RBUTTONUP, win.WM_LBUTTONUP:
		y := int(int32(lParam) >> 16)
		x := int(int32(lParam) & 0xFFFF)
		window.MouseEvent(MG_MOUSE_UP, x, y)
		break
	case win.WM_DESTROY:
		fmt.Println("-------quit--------")
		win.PostQuitMessage(0)
		syscall.Exit(0)
		break
	default:
		return win.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	return 0
}

func LoopEvent() {
	var msg win.MSG
	for {
		if win.GetMessage(&msg, 0, 0, 0) == win.TRUE {
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		} else {
			break
		}
	}
}

func initWindow(hwnd win.HWND, width, height int32) *Window {
	hdc := win.GetDC(hwnd)
	mMemdc := win.CreateCompatibleDC(hdc)
	h := win.BITMAPINFO{}
	h0 := win.BITMAPINFOHEADER{}
	h0.BiHeight = height
	h0.BiWidth = width
	h0.BiPlanes = 1
	h0.BiBitCount = 24
	h0.BiSizeImage = uint32((h0.BiWidth*int32(h0.BiBitCount/8) + 3) / 4 * 4 * h0.BiHeight)
	h0.BiSize = 40
	h.BmiHeader = h0
	w := &Window{}
	w.BmpInfo = h

	var lpBits unsafe.Pointer
	hBitmap := win.CreateDIBSection(mMemdc, &h0, win.DIB_RGB_COLORS, &lpBits, 0, 0)
	bitmapArray := (*[1 << 30]byte)(lpBits)
	line := int((width*int32(24/8) + 3) / 4 * 4)
	for h := 0; h < int(height); h = h + 1 {
		for w := 0; w < int(width*3); w = w + 3 {
			bitmapArray[h*line+w] = 255   // b
			bitmapArray[h*line+w+1] = 255 // g
			bitmapArray[h*line+w+2] = 255 // r
		}
	}


	w.Hwnd = hwnd
	w.mMemdc = mMemdc
	w.Hbitmap = hBitmap
	w.Bytes = bitmapArray
	w.hdc = hdc
	w.Width = int(width)
	w.Height = int(height)
	w.Line = int32((width*int32(24/8) + 3) / 4 * 4)
	win.SelectObject(w.mMemdc, win.HGDIOBJ(w.Hbitmap))
	return w
}

func UpdateUI(w *Window) {
	ticker := time.NewTicker(time.Millisecond * 50)
	for { //循环
		<-ticker.C
		if w.UpdateFlag {
			win.BitBlt(w.hdc, 0, 0, w.BmpInfo.BmiHeader.BiWidth, w.BmpInfo.BmiHeader.BiHeight, w.mMemdc, 0, 0, win.SRCCOPY)
		}
		w.UpdateFlag = false
	}
}
