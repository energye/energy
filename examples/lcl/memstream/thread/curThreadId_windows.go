package thread

import "github.com/energye/energy/v2/pkgs/win"

func GetCurrentThreadId() uintptr {
	return win.GetCurrentThreadId()
}
