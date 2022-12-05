//go:build go1.12 && windows

package nclient4

import (
	"errors"
	"net"
)

func NewRawUDPConn(iface string, port int) (net.PacketConn, error) {
	return nil, errors.New("not supported")
}
