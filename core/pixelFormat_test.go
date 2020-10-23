package core

import "testing"

func TestNewPixelFormat(t *testing.T) {

	bytepf := []byte{64, 20, 0, 1, 120, 120, 250, 250, 10, 10, 19, 16, 156, 3, 2, 78}
	_, err := NewPixelFormat(bytepf)

	if err != nil {
		t.Errorf("Test failed, wrong size of byte array")
	}

	_, err = NewPixelFormat([]byte{})

	if err == nil {
		t.Errorf("Test failed, PixelFormat creation should abort for wrong  argument reasons")
	}
}
