// package dialer defines the Dialer interface.
//
// The interface is used in several different places throughout the
// spiderproxy source tree, and matches interfaces used by other
// libraries as well.
package dialer // "zenhack.net/go/spiderproxy/p/dialer"

import (
	"net"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}
