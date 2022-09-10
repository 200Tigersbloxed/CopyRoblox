package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/gonutz/w32/v2"
	"os"
)

// #include <Windows.h>
import "C"

var (
	instances *systray.MenuItem
	console   w32.HWND
)

func cacheConsoleWindow() {
	consoleA := w32.GetConsoleWindow()
	console = consoleA
	fmt.Println("Found and cached ConsoleWindow")
}

func hideConsoleWindow() {
	w32.ShowWindowAsync(console, w32.SW_HIDE)
}

func main() {
	cacheConsoleWindow()
	manualStr := "ROBLOX_singletonMutex"
	mutex := C.CreateMutexA(nil, 1, C.CString(manualStr))
	fmt.Println("Started CopyRoblox")
	hideConsoleWindow()
	systray.Run(func() {
		systray.SetIcon(IconData)
		systray.SetTitle("CopyRoblox")
		systray.AddMenuItem("CopyRoblox", "")
		systray.AddSeparator()
		exitItem := systray.AddMenuItem("Exit", "This will close out of all unfocused Roblox applications!")
		go func() {
			<-exitItem.ClickedCh
			systray.Quit()
		}()
	}, func() {
		fmt.Println("Goodbye!")
		C.ReleaseMutex(mutex)
		C.CloseHandle(mutex)
		os.Exit(0)
	})
}
