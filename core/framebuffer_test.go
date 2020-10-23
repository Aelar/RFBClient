package core

import "testing"

func initPF() PixelFormat {
	bytepf := []byte{64, 20, 0, 1, 120, 120, 250, 250, 10, 10, 19, 16, 156, 3, 2, 78}
	pf, _ := NewPixelFormat(bytepf)
	return pf
}

func TestNewframebuffer(t *testing.T) {
	fbw := []byte{128, 255}
	fbh := []byte{100, 50}
	bytepf := []byte{64, 20, 0, 1, 120, 120, 250, 250, 10, 10, 19, 16, 156, 3, 2, 78}
	nl := []byte{0, 0, 0, 4}
	ns := []byte{0x74, 0x65, 0x73, 0x74}
	fb, err := NewFramebuffer(fbw, fbh, bytepf, nl, ns)

	if fb.nameString != "test" {
		t.Errorf("Bad  value assignment, expected test, got %s", fb.nameString)
	} else if err != nil {
		t.Errorf(err.Error())
	}

}
