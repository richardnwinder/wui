//+build windows

package wui

import "github.com/gonutz/w32"

func MessageBox(parent *Window, caption, text string) {
	msgBox(parent, caption, text, w32.MB_OK)
}

func MessageBoxError(parent *Window, caption, text string) {
	msgBox(parent, caption, text, w32.MB_OK|w32.MB_ICONERROR)
}

func MessageBoxWarning(parent *Window, caption, text string) {
	msgBox(parent, caption, text, w32.MB_OK|w32.MB_ICONWARNING)
}

func MessageBoxInfo(parent *Window, caption, text string) {
	msgBox(parent, caption, text, w32.MB_OK|w32.MB_ICONINFORMATION)
}

func MessageBoxQuestion(parent *Window, caption, text string) {
	msgBox(parent, caption, text, w32.MB_OK|w32.MB_ICONQUESTION)
}

func MessageBoxOKCancel(parent *Window, caption, text string) bool {
	return msgBox(parent, caption, text, w32.MB_OKCANCEL) == w32.IDOK
}

func MessageBoxYesNo(parent *Window, caption, text string) bool {
	return msgBox(parent, caption, text, w32.MB_YESNO|w32.MB_ICONQUESTION) == w32.IDYES
}

func MessageBoxCustom(parent *Window, caption, text string, flags uint) int {
	return msgBox(parent, caption, text, flags)
}

func msgBox(parent *Window, caption, text string, flags uint) int {
	var handle w32.HWND
	if parent != nil {
		handle = parent.handle
	}
	return w32.MessageBox(handle, text, caption, w32.MB_TOPMOST|flags)
}
