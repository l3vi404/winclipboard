package winclipboard

import (
	"context"
	"fmt"
	"github.com/l3vi404/winclipboard/win32"
	"golang.org/x/sys/windows"
	"sync"
	"sync/atomic"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// SetClipboardText
//
// Clipboard write text.
//
// Params:
//
//   - text: need write text.
func SetClipboardText(text string) error {
	// Convert the string to UTF-16 and add a null terminator
	utf16Text := utf16.Encode([]rune(text + "\x00"))
	size := len(utf16Text) * 2

	// Open clipboard
	r, _, err := win32.OpenClipboard.Call(0)
	if r == 0 {
		return fmt.Errorf("failed to open clipboard: %v", err)
	}
	defer win32.CloseClipboard.Call()

	// Clear clipboard
	win32.EmptyClipboard.Call()

	// Allocate global memory
	hMem, _, err := win32.GlobalAlloc.Call(uintptr(win32.GMEM_MOVEABLE), uintptr(size))
	if hMem == 0 {
		return fmt.Errorf("memory allocation failed: %v", err)
	}

	// Lock memory to write content
	ptr, _, err := win32.GlobalLock.Call(hMem)
	if ptr == 0 {
		return fmt.Errorf("failed to lock memory: %v", err)
	}
	defer win32.GlobalUnlock.Call(hMem)

	// Write UTF-16 data to memory
	dst := (*[1 << 20]uint16)(unsafe.Pointer(ptr))[:len(utf16Text)]
	copy(dst, utf16Text)

	// Set clipboard data
	r, _, err = win32.SetClipboardData.Call(uintptr(win32.CF_UNICODETEXT), hMem)
	if r == 0 {
		return fmt.Errorf("failed to set clipboard: %v", err)
	}

	return nil
}

// IsFormatAvailable
//
// Check format availability.
//
// Params:
//
//   - format: clipboard format value, example: format CF_UNICODETEXT, value is 13.
func IsFormatAvailable(format uint32) bool {
	ret, _, _ := win32.IsClipboardFormatAvailable.Call(uintptr(format))
	return ret != 0
}

// GetClipboardTextByFormat
//
// Get clipboard text by format.
//
// Params:
//
//   - format: clipboard format value, example: format CF_UNICODETEXT, value is 13.
func GetClipboardTextByFormat(format uint32) (string, error) {
	hMem, _, _ := win32.GetClipboardData.Call(uintptr(format))
	if hMem == 0 {
		return "", fmt.Errorf("GetClipboardData failed for format %d", format)
	}

	ptr, _, _ := win32.GlobalLock.Call(hMem)
	if ptr == 0 {
		return "", fmt.Errorf("GlobalLock failed for format %d", format)
	}
	defer win32.GlobalUnlock.Call(hMem)

	switch format {
	case win32.CF_UNICODETEXT:
		text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(ptr))[:])
		return text, nil
	case win32.CF_TEXT, win32.CF_OEMTEXT:
		// ANSI encoding, simple conversion (possibly garbled)
		text := ""
		for i := 0; ; i++ {
			ch := *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
			if ch == 0 {
				break
			}
			text += string(ch)
		}
		return text, nil
	default:
		return "", fmt.Errorf("unsupported format")
	}
}

// GetPreferredClipboardFormat
//
// Params:
//
//   - formats: clipboard format value array, example: format CF_UNICODETEXT, value is 13.
func GetPreferredClipboardFormat(formats []uint32) (uint32, error) {
	if len(formats) == 0 {
		return 0, fmt.Errorf("empty format list")
	}

	ptr := unsafe.Pointer(&formats[0])
	ret, _, _ := win32.GetPriorityClipboardFormat.Call(
		uintptr(ptr),
		uintptr(len(formats)),
	)

	r := int32(ret)
	if r == -1 {
		return 0, fmt.Errorf("none of the specified formats is available")
	}
	if r == -2 {
		return 0, fmt.Errorf("clipboard error occurred")
	}

	return uint32(r), nil
}

// GetClipboardText
//
// Read the latest content from the clipboard, but only text can be read. If the latest content read is not text, an error will be returned.
func GetClipboardText() (string, error) {
	// Open clipboard
	r, _, err := win32.OpenClipboard.Call(0)
	if r == 0 {
		return "", fmt.Errorf("cannot open clipboard: %v", err)
	}
	defer win32.CloseClipboard.Call()

	// Priority 1 CF_UNICODETEXT
	if IsFormatAvailable(win32.CF_UNICODETEXT) {
		return GetClipboardTextByFormat(win32.CF_UNICODETEXT)
	}
	// Priority 2 CF_TEXT
	if IsFormatAvailable(win32.CF_TEXT) {
		return GetClipboardTextByFormat(win32.CF_TEXT)
	}
	// Priority 3 CF_OEMTEXT
	if IsFormatAvailable(win32.CF_OEMTEXT) {
		return GetClipboardTextByFormat(win32.CF_OEMTEXT)
	}

	return "", fmt.Errorf("the clipboard does not contain supported text formats")
}

type ClipboardCallback func(ctx any)

type clipboardListener struct {
	handler ClipboardCallback
	context any
}

var (
	clipboardUpdateListenersMutex  sync.Mutex
	clipboardUpdateListeners       []clipboardListener
	clipboardDestroyListenersMutex sync.Mutex
	clipboardDestroyListeners      []clipboardListener
)

// AddClipboardUpdateListener
//
// Add clipboard update listener.
func AddClipboardUpdateListener(handler ClipboardCallback, ctx context.Context) {
	clipboardUpdateListenersMutex.Lock()
	defer clipboardUpdateListenersMutex.Unlock()
	clipboardUpdateListeners = append(clipboardUpdateListeners, clipboardListener{
		handler: handler,
		context: ctx,
	})
}

// ClearClipboardUpdateListeners
//
// Clear clipboard update listener.
func ClearClipboardUpdateListeners() {
	clipboardUpdateListenersMutex.Lock()
	defer clipboardUpdateListenersMutex.Unlock()
	clipboardUpdateListeners = nil
}

// AddClipboardDestroyListener
//
// Add clipboard destroy listener
func AddClipboardDestroyListener(handler ClipboardCallback, ctx context.Context) {
	clipboardDestroyListenersMutex.Lock()
	defer clipboardDestroyListenersMutex.Unlock()
	clipboardDestroyListeners = append(clipboardDestroyListeners, clipboardListener{
		handler: handler,
		context: ctx,
	})
}

// ClearClipboardDestroyListeners
//
// Clear clipboard destroy listener
func ClearClipboardDestroyListeners() {
	clipboardDestroyListenersMutex.Lock()
	defer clipboardDestroyListenersMutex.Unlock()
	clipboardDestroyListeners = nil
}

func utf16PtrFromString(s string) *uint16 {
	ptr, err := windows.UTF16PtrFromString(s)
	if err != nil {
		panic(err)
	}
	return ptr
}

// wndProc
//
// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/api/winuser/nc-winuser-wndproc
func wndProc(hwnd windows.Handle, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win32.WM_CLIPBOARDUPDATE:
		clipboardUpdateListenersMutex.Lock()
		defer clipboardUpdateListenersMutex.Unlock()
		for _, listener := range clipboardUpdateListeners {
			go listener.handler(listener.context)
		}
	case win32.WM_DESTROY:
		clipboardUpdateListenersMutex.Lock()
		clipboardDestroyListenersMutex.Lock()
		defer func() {
			ClearClipboardUpdateListeners()
			ClearClipboardDestroyListeners()
			clipboardUpdateListenersMutex.Unlock()
			clipboardDestroyListenersMutex.Unlock()
		}()
		for _, listener := range clipboardDestroyListeners {
			go listener.handler(listener.context)
		}
		win32.RemoveClipboardFormatListener.Call(uintptr(hwnd))
		win32.PostQuitMessage.Call(0)
	default:
		ret, _, _ := win32.DefWindowProcW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
		return ret
	}
	return 0
}

// createClipboardListenerHideWindow
//
// Create clipboard listener hide window.
func createClipboardListenerHideWindow() (uintptr, error) {
	var hInstance windows.Handle
	err := windows.GetModuleHandleEx(0, nil, &hInstance)
	if err != nil {
		return 0, err
	}

	className := utf16PtrFromString("ClipboardListenerWindow")
	wndClass := win32.WNDCLASSEXW{
		CbSize:        uint32(unsafe.Sizeof(win32.WNDCLASSEXW{})),
		Style:         win32.CS_HREDRAW | win32.CS_VREDRAW,
		LpfnWndProc:   syscall.NewCallback(wndProc),
		HInstance:     hInstance,
		LpszClassName: className,
	}

	atom, _, err := win32.RegisterClassExW.Call(uintptr(unsafe.Pointer(&wndClass)))
	if atom == 0 {
		return 0, err
	}

	hwnd, _, err := win32.CreateWindowExW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("ClipboardListenerWindow"))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("ClipboardListenerHideWindow"))),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(hInstance),
		uintptr(0),
	)
	if hwnd == 0 {
		return 0, err
	}
	return hwnd, nil
}

var (
	clipboardListenerHwnd uintptr
	clipboardLoopRunning  atomic.Bool
)

// StartClipboardListener
//
// Start clipboard listener
func StartClipboardListener() error {
	if clipboardLoopRunning.Load() {
		return fmt.Errorf("clipboard listener already running")
	}
	hwnd, err := createClipboardListenerHideWindow()
	if hwnd == 0 {
		return fmt.Errorf("startClipboardListener errer: %w", err)
	}
	clipboardListenerHwnd = hwnd
	clipboardLoopRunning.Store(true)
	win32.AddClipboardFormatListener.Call(hwnd)

	var msg win32.MSG
	for {
		ret, _, _ := win32.GetMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if int32(ret) == 0 {
			break
		}
		win32.TranslateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		win32.DispatchMessageW.Call(uintptr(unsafe.Pointer(&msg)))
	}
	clipboardLoopRunning.Store(false)
	return nil
}

// StopClipboardListener
//
// Stop clipboard Listener.
func StopClipboardListener() {
	if clipboardListenerHwnd != 0 && clipboardLoopRunning.Load() {
		win32.PostMessageW.Call(clipboardListenerHwnd, win32.WM_DESTROY, 0, 0)
	}
}
