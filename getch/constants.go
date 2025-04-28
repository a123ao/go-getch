package getch

const (
	// STD_INPUT_HANDLE is the standard input handle
	STD_INPUT_HANDLE = uint32(-10 & 0xFFFFFFFF)

	// Console modes
	ENABLE_LINE_INPUT      = uint32(0x0002)
	ENABLE_ECHO_INPUT      = uint32(0x0004)
	ENABLE_PROCESSED_INPUT = uint32(0x0001)

	// INPUT_RECORD types
	KEY_EVENT = uint16(0x0001) // Keyboard event

	// VIRTUAL KEYS
	VK_UP    = uint16(0x26) // Up arrow key
	VK_DOWN  = uint16(0x28) // Down arrow key
	VK_LEFT  = uint16(0x25) // Left arrow key
	VK_RIGHT = uint16(0x27) // Right arrow key
	VK_ENTER = uint16(0x0D) // Enter key
	VK_SHIFT = uint16(0x10) // Shift key
)
