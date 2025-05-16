package win32

import "golang.org/x/sys/windows"

var (
	user32 = windows.NewLazySystemDLL("user32.dll")

	// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-createwindowexw

	CreateWindowExW = user32.NewProc("CreateWindowExW")

	// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getmessagew

	GetMessageW = user32.NewProc("GetMessageW")

	// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/_dataxchg/, in Winuser.h menu

	AddClipboardFormatListener    = user32.NewProc("AddClipboardFormatListener")
	CloseClipboard                = user32.NewProc("CloseClipboard")
	EmptyClipboard                = user32.NewProc("EmptyClipboard")
	EnumClipboardFormats          = user32.NewProc("EnumClipboardFormats")
	GetClipboardData              = user32.NewProc("GetClipboardData")
	GetClipboardFormatNameW       = user32.NewProc("GetClipboardFormatNameW")
	GetPriorityClipboardFormat    = user32.NewProc("GetPriorityClipboardFormat")
	IsClipboardFormatAvailable    = user32.NewProc("IsClipboardFormatAvailable")
	OpenClipboard                 = user32.NewProc("OpenClipboard")
	RemoveClipboardFormatListener = user32.NewProc("RemoveClipboardFormatListener")
	SetClipboardData              = user32.NewProc("SetClipboardData")
	RegisterClassExW              = user32.NewProc("RegisterClassExW")
	PostQuitMessage               = user32.NewProc("PostQuitMessage")
	DefWindowProcW                = user32.NewProc("DefWindowProcW")
	TranslateMessage              = user32.NewProc("TranslateMessage")
	DispatchMessageW              = user32.NewProc("DispatchMessageW")
	PostMessageW                  = user32.NewProc("PostMessageW")
)

type WNDCLASSEXW struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     windows.Handle
	HIcon         windows.Handle
	HCursor       windows.Handle
	HbrBackground windows.Handle
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       windows.Handle
}

// POINT
//
// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/gdiplustypes/nl-gdiplustypes-point
type POINT struct {
	X, Y int32
}

// MSG
//
// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getmessage
type MSG struct {
	Hwnd    windows.Handle
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}
