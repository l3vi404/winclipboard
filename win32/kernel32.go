package win32

import "golang.org/x/sys/windows"

// https://learn.microsoft.com/zh-cn/windows/win32/api/winbase/
var (
	kernel32 = windows.NewLazySystemDLL("kernel32")

	GlobalAlloc  = kernel32.NewProc("GlobalAlloc")
	GlobalLock   = kernel32.NewProc("GlobalLock")
	GlobalUnlock = kernel32.NewProc("GlobalUnlock")
)
