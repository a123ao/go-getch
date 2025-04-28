package getch

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	isInitialized bool
	hstdin        windows.Handle
	consoleMode   uint32
)

func initialize() error {
	// Get the standard input handle
	handle, _, _ := procGetStdHandle.Call(uintptr(STD_INPUT_HANDLE))
	if handle == 0 {
		return fmt.Errorf("failed to get the standard input handle: %w", windows.GetLastError())
	}
	hstdin = windows.Handle(handle)

	// Get console mode
	ret, _, _ := procGetConsoleMode.Call(uintptr(hstdin), uintptr(unsafe.Pointer(&consoleMode)))
	if ret == 0 {
		return fmt.Errorf("failed to get console mode: %w", windows.GetLastError())
	}

	isInitialized = true
	return nil
}

// Get a single character from the standard input without echoing it to the console.
func Getch() (KEY, error) {
	if !isInitialized {
		if err := initialize(); err != nil {
			return KEY{}, fmt.Errorf("failed to initialize: %w", err)
		}
	}

	var ret uintptr

	// Change console mode to disable echo and line input
	newConsoleMode := consoleMode &^ (ENABLE_LINE_INPUT | ENABLE_ECHO_INPUT)
	ret, _, _ = procSetConsoleMode.Call(uintptr(hstdin), uintptr(newConsoleMode))
	if ret == 0 {
		return KEY{}, fmt.Errorf("failed to set console mode: %w", windows.GetLastError())
	}

	// Defer restoring the console mode
	defer func() {
		procSetConsoleMode.Call(uintptr(hstdin), uintptr(consoleMode))
	}()

	// Read a single character
	var inputRecord INPUT_RECORD
	var eventsRead uint32

	for {
		// Read console input
		ret, _, _ := procReadConsoleInput.Call(
			uintptr(hstdin),
			uintptr(unsafe.Pointer(&inputRecord)),
			1, // Number of records to read
			uintptr(unsafe.Pointer(&eventsRead)),
		)

		if ret == 0 {
			return KEY{}, fmt.Errorf("failed to read console input: %w", windows.GetLastError())
		}

		if eventsRead == 0 {
			continue
		}

		if inputRecord.EventType != KEY_EVENT {
			continue
		}

		// Get the key event record from the union field
		keyEvent := (*KEY_EVENT_RECORD)(unsafe.Pointer(&inputRecord.Event[0]))

		if keyEvent.KeyDown != 0 {
			return KEY{
				UnicodeChar:    keyEvent.UnicodeChar,
				VirtualKeyCode: keyEvent.VirtualKeyCode,
			}, nil
		}
	}
}
