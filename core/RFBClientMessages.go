package core

import "strconv"

//SetPixelFormat request Client-to-server message
type SetPixelFormat struct {
	messageType uint8
	//3 bytes
	padding     []byte
	pixelformat PixelFormat
}

//NewSetPixelFormat Create a new SetPixelFormat ready to be sent
func NewSetPixelFormat(padding []byte, px PixelFormat) (SetPixelFormat, error) {

	if len(padding) != 3 {
		return SetPixelFormat{}, &FormatError{string("Expecting [3]byte, got " + strconv.Itoa(len(padding))), -1}
	}
	return SetPixelFormat{
		uint8(0),
		padding,
		px,
	}, nil

}

//SetEncoding request client-to-server message
//Should  befolowed  by number-of-encoding message(s) of type rune(alias for int32)
type SetEncoding struct {
	messageType       uint8
	padding           uint8
	numberOfEncodings uint16
}

//NewSetEncoding Create a new SetEncoding ready to be sent
func NewSetEncoding(padding uint8, numberOfEncodings uint16) (SetEncoding, error) {

	se := SetEncoding{2, padding, numberOfEncodings}

	return se, nil

}

//GetNumberOfEncoding return the mapping  of  encodings, ready for sending to server via rfb socket
func GetNumberOfEncoding() map[string]int32 {

	return map[string]int32{
		"RAW":                         0,
		"CopyRect":                    1,
		"RRE":                         2,
		"Hextile":                     5,
		"TRLE":                        15,
		"ZRLE":                        16,
		"Cursor pseudo-encoding":      -239,
		"DesktopSize pseudo-encoding": -223,
	}

}

//FramebufferUpdateRequest asking for update on a specific area of the screen
type FramebufferUpdateRequest struct {
	messageType uint8
	incremental uint8
	xPosition   uint16
	yPosition   uint16
	width       uint16
	height      uint16
}

//NewFramebufferUpdateRequest return a new request for a an update  on a  specific area of screen
func NewFramebufferUpdateRequest(incremental uint8,
	xPosition uint16,
	yPosition uint16,
	width uint16,
	height uint16) (FramebufferUpdateRequest, error) {

	return FramebufferUpdateRequest{3, incremental, xPosition, yPosition, width, height}, nil
}
