package win32

// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/dataxchg/clipboard-notifications

const (
	GMEM_MOVEABLE uint32 = 0x0002
)

// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/dataxchg/clipboard-notifications

const (
	WM_CLIPBOARDUPDATE uint32 = 0x031D
)

// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/winmsg/window-notifications

const (
	WM_DESTROY        = 0x0002
	WM_QUIT    uint32 = 0x0012
)

// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/dataxchg/standard-clipboard-formats

const (
	CF_BITMAP          uint32 = 2
	CF_DIB             uint32 = 8
	CF_DIBV5           uint32 = 17
	CF_DIF             uint32 = 5
	CF_DSPBITMAP       uint32 = 0x0082
	CF_DSPENHMETAFILE  uint32 = 0x008E
	CF_DSPMETAFILEPICT uint32 = 0x0083
	CF_DSPTEXT         uint32 = 0x0081
	CF_ENHMETAFILE     uint32 = 14
	CF_GDIOBJFIRST     uint32 = 0x0300
	CF_GDIOBJLAST      uint32 = 0x03FF
	CF_HDROP           uint32 = 15
	CF_LOCALE          uint32 = 16
	CF_METAFILEPICT    uint32 = 3
	CF_OEMTEXT         uint32 = 7
	CF_OWNERDISPLAY    uint32 = 0x0080
	CF_PALETTE         uint32 = 9
	CF_PENDATA         uint32 = 10
	CF_PRIVATEFIRST    uint32 = 0x0200
	CF_PRIVATELAST     uint32 = 0x02FF
	CF_RIFF            uint32 = 11
	CF_SYLK            uint32 = 4
	CF_TEXT            uint32 = 1
	CF_TIFF            uint32 = 6
	CF_UNICODETEXT     uint32 = 13
	CF_WAVE            uint32 = 12
)

// Reference resources: https://learn.microsoft.com/zh-cn/windows/win32/winmsg/window-class-styles

const (
	CS_VREDRAW = 0x0001
	CS_HREDRAW = 0x0002
)
