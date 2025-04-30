package process

import "golang.org/x/sys/windows"

var (
	Kernel32DLL = windows.NewLazyDLL("kernel32.dll")

	// procGetStdHandle is a handle to the standard IO
	ProcGetStdHandle = Kernel32DLL.NewProc("GetStdHandle")

	// procGetConsoleMode is a handle to the console mode
	ProcGetConsoleMode = Kernel32DLL.NewProc("GetConsoleMode")

	// procSetConsoleMode is a handle to set the console mode
	ProcSetConsoleMode = Kernel32DLL.NewProc("SetConsoleMode")

	// procReadConsole is a handle to read the console
	ProcReadConsoleInput = Kernel32DLL.NewProc("ReadConsoleInputW")
)
