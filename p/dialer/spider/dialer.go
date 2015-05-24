package spider

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"zenhack.net/go/spiderproxy/p/dialer"
	"zenhack.net/go/spiderproxy/p/dialer/glob"
)

func NewDialer(node *Node, root dialer.Dialer, agent agent.Agent) (dialer.Dialer, error) {
	clientConfig := &ssh.ClientConfig{
		User: node.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.Signers),
		},
	}
	globDialer := glob.NewDialer(root)

	// If we just do build := ..., then the variable build won't be visible
	// inside the function. This is a hack to get around that:
	var build func(node *Node, d dialer.Dialer) error
	build = func(node *Node, d dialer.Dialer) error {
		address := net.JoinHostPort(node.Host, fmt.Sprint(node.Port))
		conn, err := d.Dial("tcp", address)
		if err != nil {
			return err
		}
		sshConn, chans, reqs, err := ssh.NewClientConn(conn, address, clientConfig)
		if err != nil {
			conn.Close()
			return err
		}
		client := ssh.NewClient(sshConn, chans, reqs)
		for i := range(node.Match) {
			globDialer.Append(node.Match[i], client)
		}
		for i := range(node.Next) {
			err := build(node.Next[i], client)
			if err != nil {
				sshConn.Close()
				conn.Close()
				return err
			}
		}
		return nil
	}
	err := build(node, root)
	if err != nil {
		return nil, err
	}
	return globDialer, nil
}
