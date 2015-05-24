package glob

import (
	"net"
	"path/filepath" // for globbing

	"zenhack.net/go/spiderproxy/p/dialer"
)

type Dialer struct {
	globs []string
	dialers []dialer.Dialer
	fallback dialer.Dialer
}

func (d *Dialer) Append(pattern string, dialer dialer.Dialer) {
	d.globs = append(d.globs, pattern)
	d.dialers = append(d.dialers, dialer)
}

func NewDialer(fallback dialer.Dialer) *Dialer {
	return &Dialer {
		globs: []string{},
		dialers: []dialer.Dialer{},
		fallback: fallback,
	}
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	for i := range(d.globs) {
		if match, _ := filepath.Match(d.globs[i], address); match {
			return d.dialers[i].Dial(network, address)
		}
	}
	return d.fallback.Dial(network, address)
}
