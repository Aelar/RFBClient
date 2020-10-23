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
type SetEncoding struct {
	messageType       uint8
	padding           uint8
	numberOfEncodings uint16
}

func NewSetEncoding(padding uint8, numberOfEncodings uint16) (SetEncoding, error) {

}
