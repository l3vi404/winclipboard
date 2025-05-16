package win32

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

func TestEnumClipboardFormats(t *testing.T) {
	r, _, err := OpenClipboard.Call(0)
	if r == 0 {
		t.Errorf("OpenClipboard failed: %v", err)
		return
	}
	defer CloseClipboard.Call()

	var format uint32 = 0
	for {
		ret, _, _ := EnumClipboardFormats.Call(uintptr(format))
		if ret == 0 {
			break
		}

		format = uint32(ret)

		var nameBuf [256]uint16
		nameLen, _, _ := GetClipboardFormatNameW.Call(
			uintptr(format),
			uintptr(unsafe.Pointer(&nameBuf[0])),
			uintptr(len(nameBuf)),
		)

		if nameLen > 0 {
			name := syscall.UTF16ToString(nameBuf[:nameLen])
			fmt.Printf("Format: %d, Name: %s\n", format, name)
		} else {
			fmt.Printf("Format: %d (predefined or unnamed)\n", format)
		}
	}
}
