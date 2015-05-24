package sshutil

import (
	"golang.org/x/crypto/ssh/agent"

	"io"
	"net"
	"os"
)

type agentCloser struct {
	agent.Agent
	conn net.Conn
}

type AgentCloser interface {
	agent.Agent
	io.Closer
}

func (ac *agentCloser) Close() error {
	return ac.conn.Close()
}

func DialAgent() (AgentCloser, error) {
	authSock := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", authSock)
	if err != nil {
		return nil, err
	}
	return &agentCloser{agent.NewClient(conn), conn}, nil
}
