package win32

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Printf("WM_CLIPBOARDUPDATE: %d\n", WM_CLIPBOARDUPDATE)
	fmt.Printf("WM_QUIT: %d\n", WM_QUIT)

	fmt.Printf("CF_BITMAP: %d\n", CF_BITMAP)
	fmt.Printf("CF_DIB: %d\n", CF_DIB)
	fmt.Printf("CF_DIBV5: %d\n", CF_DIBV5)
	fmt.Printf("CF_DIF: %d\n", CF_DIF)
	fmt.Printf("CF_DSPBITMAP: %d\n", CF_DSPBITMAP)
	fmt.Printf("CF_DSPENHMETAFILE: %d\n", CF_DSPENHMETAFILE)
	fmt.Printf("CF_DSPMETAFILEPICT: %d\n", CF_DSPMETAFILEPICT)
	fmt.Printf("CF_DSPTEXT: %d\n", CF_DSPTEXT)
	fmt.Printf("CF_ENHMETAFILE: %d\n", CF_ENHMETAFILE)
	fmt.Printf("CF_GDIOBJFIRST: %d\n", CF_GDIOBJFIRST)
	fmt.Printf("CF_GDIOBJLAST: %d\n", CF_GDIOBJLAST)
	fmt.Printf("CF_HDROP: %d\n", CF_HDROP)
	fmt.Printf("CF_LOCALE: %d\n", CF_LOCALE)
	fmt.Printf("CF_METAFILEPICT: %d\n", CF_METAFILEPICT)
	fmt.Printf("CF_OWNERDISPLAY: %d\n", CF_OWNERDISPLAY)
	fmt.Printf("CF_PALETTE: %d\n", CF_PALETTE)
	fmt.Printf("CF_PENDATA: %d\n", CF_PENDATA)
	fmt.Printf("CF_PRIVATEFIRST: %d\n", CF_PRIVATEFIRST)
	fmt.Printf("CF_PRIVATELAST: %d\n", CF_PRIVATELAST)
	fmt.Printf("CF_RIFF: %d\n", CF_RIFF)
	fmt.Printf("CF_SYLK: %d\n", CF_SYLK)
	fmt.Printf("CF_TIFF: %d\n", CF_TIFF)
	fmt.Printf("CF_UNICODETEXT: %d\n", CF_UNICODETEXT)
	fmt.Printf("CF_WAVE: %d\n", CF_WAVE)

	fmt.Printf("CS_VREDRAW: %d\n", CS_VREDRAW)
	fmt.Printf("CS_HREDRAW: %d\n", CS_HREDRAW)
}
