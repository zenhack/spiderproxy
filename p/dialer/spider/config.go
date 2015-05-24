package spider

import (
	"encoding/json"
	"io"
	"os"
)

type Config []*Node

type Node struct {
	Comment string   `json:"comment,omitempty"`
	User    string   `json:"user,omitempty"`
	Host    string   `json:"host"`
	Port    uint16   `json:"port,omitempty"`
	Match   []string `json:"match"`
	Next    []*Node  `json:"next,omitempty"`
}

func LoadConfig(r io.Reader) (Config, error) {
	var ret []*Node
	dec := json.NewDecoder(r)
	err := dec.Decode(&ret)
	if err == nil {
		for i := range(ret) {
			ret[i].normalize()
		}
	}
	return ret, err
}

func (n *Node) normalize() {
	if n.User == "" { n.User = os.Getenv("USER") }
	if n.Port == 0 { n.Port = 22 }
	for i := range(n.Next) {
		n.Next[i].normalize()
	}
}
