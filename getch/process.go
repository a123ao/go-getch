package getch

import "golang.org/x/sys/windows"

var (
	kernel32DLL = windows.NewLazyDLL("kernel32.dll")

	// procGetStdHandle is a handle to the standard IO
	procGetStdHandle = kernel32DLL.NewProc("GetStdHandle")

	// procGetConsoleMode is a handle to the console mode
	procGetConsoleMode = kernel32DLL.NewProc("GetConsoleMode")

	// procSetConsoleMode is a handle to set the console mode
	procSetConsoleMode = kernel32DLL.NewProc("SetConsoleMode")

	// procReadConsole is a handle to read the console
	procReadConsoleInput = kernel32DLL.NewProc("ReadConsoleInputW")
)
