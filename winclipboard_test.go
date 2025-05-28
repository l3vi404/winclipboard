package winclipboard

import (
	"context"
	"fmt"
	"github.com/l3vi404/winclipboard/win32"
	"testing"
	"time"
)

func TestSetClipboardText(t *testing.T) {
	err := SetClipboardText("TestSetClipboardText")
	if err != nil {
		t.Errorf("Faild SetClipboardText: %v", err)
	}
}

func TestIsFormatAvailable(t *testing.T) {
	err := SetClipboardText("TestSetClipboardText")
	if err != nil {
		t.Errorf("Faild SetClipboardText: %v", err)
	}
	IsFormatAvailable(win32.CF_UNICODETEXT)
}

func TestGetClipboardText(t *testing.T) {
	err := SetClipboardText("TestGetClipboardText")
	if err != nil {
		t.Errorf("Faild SetClipboardText: %v", err)
	}
	_, err = GetClipboardText()
	if err != nil {
		t.Errorf("Faild GetClipboardText: %v", err)
	}
}

func TestGetPreferredClipboardFormat(t *testing.T) {
	formats := []uint32{win32.CF_UNICODETEXT, win32.CF_TEXT, win32.CF_OEMTEXT}
	format, err := GetPreferredClipboardFormat(formats)
	if err != nil {
		t.Errorf("No preferred format available: %v", err)
	} else {
		fmt.Printf("Preferred format is: %d\n", format)
	}
}

func TestStartClipboardListener(t *testing.T) {
	AddClipboardUpdateListener(func(ctx any) {
		text, _ := GetClipboardText()
		fmt.Printf("The clipboard has been changed, content is: %q\n", text)
	}, context.Background())

	AddClipboardDestroyListener(func(ctx any) {
		fmt.Println("The clipboard destroyed")
	}, context.Background())
	go func() {
		if err := StartClipboardListener(); err != nil {
			t.Errorf("Faild StartClipboardListener: %v", err)
		}
	}()

	time.Sleep(5 * time.Second)
	StopClipboardListener()
}
