package keyboard

// INPUT_RECORD structure represents an input event in the console input buffer
// https://learn.microsoft.com/en-us/windows/console/input-record-str
// Use explicit layout if needed, but Go usually aligns correctly for simple cases
type INPUT_RECORD struct {
	EventType uint16
	_         uint16   // Padding
	Event     [16]byte // Union field, KEY_EVENT_RECORD is the largest relevant part
}

// KEY_EVENT_RECORD structure describes a keyboard input event
// https://learn.microsoft.com/en-us/windows/console/key-event-record-str
type KEY_EVENT_RECORD struct {
	KeyDown         int32 // Using int32 for BOOL as per Go's windows package convention
	RepeatCount     uint16
	VirtualKeyCode  uint16
	VirtualScanCode uint16
	UnicodeChar     uint16 // This holds the actual character
	ControlKeyState uint32
}

// KEY structure represents a key event with Unicode character and virtual key code
type KEY_ENTRY struct {
	UnicodeChar    uint16 // Unicode character
	VirtualKeyCode uint16 // Virtual key code
}
