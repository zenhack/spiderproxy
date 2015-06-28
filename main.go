package main

import (
	"golang.org/x/crypto/ssh/agent"
	"zenhack.net/go/socks5"

	"zenhack.net/go/spiderproxy/p/dialer/spider"

	"flag"
)

import (
	"fmt"
	"net"
	"os"
)

var (
	addr        = flag.String("addr", ":1080", "port to start a socks5 server on")
	config_file = flag.String("config", "spiderproxy.json", "config file")
)

func checkfatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	file, err := os.Open(*config_file)
	checkfatal(err)
	config, err := spider.LoadConfig(file)
	checkfatal(err)

	authSock := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", authSock)
	checkfatal(err)

	dialer, err := spider.NewDialer(config, &net.Dialer{}, agent.NewClient(conn))
	checkfatal(err)
	checkfatal(socks5.ListenAndServe(dialer, *addr))
}
