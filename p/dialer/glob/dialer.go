// package glob defines a dialer which multiplexes other dialers, dispatching
// base on shell glob patterns.
package glob // import "zenhack.net/go/spiderproxy/p/dialer/glob"

import (
	"net"
	"path/filepath" // for globbing

	"zenhack.net/go/spiderproxy/p/dialer"
)

// Dialer which dispatches to other dialers based on shell glob patterns.
// The Append method can be used to add rules to the glob dialer; if no rule
// matches, a fallback (chosen upon creation, see NewDialer) is used.
type Dialer struct {
	globs    []string
	dialers  []dialer.Dialer
	fallback dialer.Dialer
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	for i := range d.globs {
		if match, _ := filepath.Match(d.globs[i], address); match {
			return d.dialers[i].Dial(network, address)
		}
	}
	return d.fallback.Dial(network, address)
}

// Add a dialer to to the glob dialer; if no previous rule matches an address,
// and "pattern" matches, then dialer will be used to make a connection.
func (d *Dialer) Append(pattern string, dialer dialer.Dialer) {
	d.globs = append(d.globs, pattern)
	d.dialers = append(d.dialers, dialer)
}

// Create a new glob dialer. The dialer will initially have no rules associated
// with it, and will dispatch all connections to "fallback".
func NewDialer(fallback dialer.Dialer) *Dialer {
	return &Dialer{
		globs:    []string{},
		dialers:  []dialer.Dialer{},
		fallback: fallback,
	}
}
