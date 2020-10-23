package core

import (
	"encoding/binary"
)

//PixelFormat 16 bytes
type PixelFormat struct {
	bitPerPixel       uint8
	depth             uint8
	bigEndianEncoding uint8
	trueColorFlag     uint8
	redMax            uint16
	greenMax          uint16
	blueMax           uint16
	redShift          uint8
	greenShift        uint8
	blueShift         uint8
	//3 bytes
	padding []byte
}

//NewPixelFormat Create a new instance of the pixel format object
func NewPixelFormat(p []byte) (PixelFormat, error) {

	if len(p) != 16 {
		return PixelFormat{}, &CustomError{"Received pixel format as wrong len"}
	}

	pf := PixelFormat{
		uint8(p[0]),
		uint8(p[1]),
		uint8(p[2]),
		uint8(p[3]),
		binary.BigEndian.Uint16(p[4:6]),
		binary.BigEndian.Uint16(p[6:8]),
		binary.BigEndian.Uint16(p[8:10]),
		uint8(p[10]),
		uint8(p[11]),
		uint8(p[12]),
		p[13:],
	}

	return pf, nil

}
