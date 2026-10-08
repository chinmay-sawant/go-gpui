package tv

import (
	"errors"
	"net"
)

// Magic builds a Wake-on-LAN packet for mac.
func Magic(mac string) ([]byte, error) {
	hw, err := net.ParseMAC(mac)
	if err != nil || len(hw) != 6 {
		return nil, errors.New("the saved TV address is unusable")
	}

	pkt := make([]byte, 102)
	for i := 0; i < 6; i++ {
		pkt[i] = 0xff
	}

	for i := 1; i <= 16; i++ {
		copy(pkt[i*6:], hw)
	}

	return pkt, nil
}

// Wake sends the magic packet on the local broadcast.
func Wake(mac string) error {
	pkt, err := Magic(mac)
	if err != nil {
		return err
	}

	conn, err := net.Dial("udp4", "255.255.255.255:9")
	if err != nil {
		return errors.New("the wake packet could not be sent")
	}

	defer conn.Close()

	_, err = conn.Write(pkt)
	if err != nil {
		return errors.New("the wake packet could not be sent")
	}

	return nil
}
