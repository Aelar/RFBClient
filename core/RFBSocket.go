package core

import (
	"bufio"
	"bytes"
	"crypto/des"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
)

//RFBSocket socket that can initialize it's  connection to rfb protocol
type RFBSocket struct {
	socket net.Conn
}

//NewRFBSocket initialize socket and connect it ot rfb remote interface
func NewRFBSocket(addr string, port string) RFBSocket {
	sock, err := net.Dial("tcp", string(addr+":"+port))

	if err != nil {
		log.Fatal(err)
	}
	rfb := RFBSocket{sock}
	return rfb
}

func (r RFBSocket) zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (r RFBSocket) desEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	src = r.zeroPadding(src, bs)
	// src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func (r RFBSocket) read(size int) []byte {
	buf := make([]byte, size)
	_, err := r.socket.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func (r RFBSocket) security(secFlag byte) error {

	if 1 == secFlag {
		r.socket.Write([]byte{1})
		return nil
	} else if 2 == secFlag {
		r.socket.Write([]byte{2})

		buf := r.read(16)
		//* TODO: Handle password prompt
		encryptedMsg, err := r.desEncrypt(buf, []byte("test"))
		if err != nil {
			return err
		}

		r.socket.Write(encryptedMsg)
		return nil
	} else {
		return &MissingPwdError{"RFB require a password", 403}
	}

}

//Handshake handshake phase for rfb protocol
func (r RFBSocket) Handshake() error {

	//Listen for handshake ex: RFB 003.008\n
	status, err := bufio.NewReader(r.socket).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//Answer to handshake ex:RFB 003.008\n
	fmt.Fprintf(r.socket, status)

	//Listen for number of security type
	buf := r.read(1)

	//Listen for auth types ex:[1]
	/*
		+--------+--------------------+
		| Number | Name               |
		+--------+--------------------+
		| 0      | Invalid            |
		| 1      | None               |
		| 2      | VNC Authentication |
		+--------+--------------------+
	*/
	nAuthTypes := int(buf[0])
	buf = r.read(nAuthTypes)

	//fmt.Println(buf)

	//Handle security type
	err = r.security(buf[0])

	if err != nil {
		log.Fatalf(err.Error())
	}

	//Listen for  success of security check
	buf = r.read(4)

	//fmt.Println(buf)

	//Send  client  init

	r.socket.Write([]byte{1})

	//receive server init

	width := r.read(2)
	height := r.read(2)
	pixelformat := r.read(16)
	nameLenght := r.read(4)
	nameString := r.read(int(binary.BigEndian.Uint32(nameLenght)))

	fb, err := NewFramebuffer(width, height, pixelformat, nameLenght, nameString)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(fb.nameString)

	return nil

}

/*func (r RFBSocket) framebufferUpdateRequest() {
	r.socket.Write([]byte{3})
}*/
