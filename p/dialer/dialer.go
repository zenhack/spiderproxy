package dialer

import (
	"net"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}
