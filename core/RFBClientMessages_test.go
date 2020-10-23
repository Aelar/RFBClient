package core

import "testing"

func TestNewSetPixelFormat(t *testing.T) {

	bytepf := []byte{64, 20, 0, 1, 120, 120, 250, 250, 10, 10, 19, 16, 156, 3, 2, 78}
	pf, _ := NewPixelFormat(bytepf)

	spf, err := NewSetPixelFormat([]byte{0, 25, 13}, pf)

	if spf.messageType != 0 {
		t.Errorf("Test failed, Wrong messageType")
	}

	if err != nil {
		t.Errorf("Test failed, Wrong padding")

	}

	_, err = NewSetPixelFormat([]byte{0, 25}, pf)

	if err == nil {
		t.Errorf("Test failed, SetPixelFormat should fail, bad  padding")

	}

}

func TestNewSetEncoding(t *testing.T) {

	se, err := NewSetEncoding([]byte{2}, []byte{2, 23})
	if spf.messageType != 0 {
		t.Errorf("Test failed, Wrong messageType")
	}
}
