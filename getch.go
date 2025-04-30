package getch

import (
	"fmt"
	"unsafe"

	"github.com/a123ao/go-getch/internal/keyboard"
	"github.com/a123ao/go-getch/internal/process"
	"golang.org/x/sys/windows"
)

var (
	// isInitialized bool
	hstdin      windows.Handle
	consoleMode uint32
)

const (
	// VIRTUAL KEYS
	VK_UP    = uint16(0x26) // Up arrow key
	VK_DOWN  = uint16(0x28) // Down arrow key
	VK_LEFT  = uint16(0x25) // Left arrow key
	VK_RIGHT = uint16(0x27) // Right arrow key
	VK_ENTER = uint16(0x0D) // Enter key
	VK_SHIFT = uint16(0x10) // Shift key
)

func setup() error {
	// Get the standard input handle
	handle, _, _ := process.ProcGetStdHandle.Call(uintptr(process.STD_INPUT_HANDLE))
	if handle == 0 {
		return fmt.Errorf("failed to get the standard input handle: %w", windows.GetLastError())
	}
	hstdin = windows.Handle(handle)

	// Get console mode
	ret, _, _ := process.ProcGetConsoleMode.Call(
		uintptr(hstdin),
		uintptr(unsafe.Pointer(&consoleMode)),
	)
	if ret == 0 {
		return fmt.Errorf("failed to get console mode: %w", windows.GetLastError())
	}

	return nil
}

func Read() (keyboard.KEY_ENTRY, error) {
	if hstdin == 0 {
		if err := setup(); err != nil {
			return keyboard.KEY_ENTRY{}, fmt.Errorf("failed to setup: %w", err)
		}
	}

	var ret uintptr

	// Set console mode to disable echo and line input
	// newConsoleMode is the current console mode with line input and echo input disabled
	newConsoleMode := consoleMode &^ (process.ENABLE_LINE_INPUT | process.ENABLE_ECHO_INPUT)
	ret, _, _ = process.ProcSetConsoleMode.Call(uintptr(hstdin), uintptr(newConsoleMode))
	if ret == 0 {
		return keyboard.KEY_ENTRY{}, fmt.Errorf("failed to set console mode: %w", windows.GetLastError())
	}

	// Defer restoring the console mode
	defer func() {
		process.ProcSetConsoleMode.Call(uintptr(hstdin), uintptr(consoleMode))
	}()

	// Read a single character
	var inputRecord keyboard.INPUT_RECORD
	var eventsRead uint32

	for {
		// Read console input
		ret, _, _ := process.ProcReadConsoleInput.Call(
			uintptr(hstdin),
			uintptr(unsafe.Pointer(&inputRecord)),
			1, // Number of records to read
			uintptr(unsafe.Pointer(&eventsRead)),
		)
		if ret == 0 {
			return keyboard.KEY_ENTRY{}, fmt.Errorf("failed to read console input: %w", windows.GetLastError())
		}

		// Check if any events were read
		if eventsRead == 0 {
			continue
		}

		// Check if the event is a key event
		if inputRecord.EventType != process.KEY_EVENT {
			continue
		}

		// Get the key event record from the union field
		keyEvent := (*keyboard.KEY_EVENT_RECORD)(unsafe.Pointer(&inputRecord.Event[0]))

		if keyEvent.KeyDown != 0 {
			return keyboard.KEY_ENTRY{
				UnicodeChar:    keyEvent.UnicodeChar,
				VirtualKeyCode: keyEvent.VirtualKeyCode,
			}, nil
		}
	}
}
