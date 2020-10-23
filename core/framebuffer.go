package core

import "encoding/binary"

//Framebuffer min 25 bytes
type Framebuffer struct {
	fbWidth     uint16
	fbHeight    uint16
	pixelFormat PixelFormat
	nameLenght  uint32
	nameString  string
}

//NewFramebuffer : Create  a new instance of  the framebuffer model
func NewFramebuffer(fbWidth []byte,
	fbHeight []byte,
	pixelFormat []byte,
	nameLenght []byte,
	nameString []byte) (Framebuffer, error) {

	pf, err := NewPixelFormat(pixelFormat)

	if err != nil {
		return Framebuffer{}, err
	}

	fb := Framebuffer{
		binary.BigEndian.Uint16(fbWidth),
		binary.BigEndian.Uint16(fbHeight),
		pf,
		binary.BigEndian.Uint32(nameLenght),
		string(nameString),
	}

	return fb, nil
}
