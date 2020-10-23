package main

import (
	core "rfbclient/core"
)

func main() {
	t := core.NewRFBSocket("127.0.0.1", "5900")

	t.Handshake()

}
