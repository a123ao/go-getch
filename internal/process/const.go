package process

const (
	// STD_INPUT_HANDLE is the standard input handle
	STD_INPUT_HANDLE = uint32(-10 & 0xFFFFFFFF)

	// Console modes
	ENABLE_LINE_INPUT      = uint32(0x0002)
	ENABLE_ECHO_INPUT      = uint32(0x0004)
	ENABLE_PROCESSED_INPUT = uint32(0x0001)

	// INPUT_RECORD types
	KEY_EVENT = uint16(0x0001) // Keyboard event
)
