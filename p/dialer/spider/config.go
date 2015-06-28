package spider

import (
	"encoding/json"
	"io"
	"os"
)

// Configuration, defining the topology the spider proxy is to implement.
type Config []*Node

// A node definition in the configuration
type Node struct {
	Comment string   `json:"comment,omitempty"`
	User    string   `json:"user,omitempty"`
	Host    string   `json:"host"`
	Port    uint16   `json:"port,omitempty"`
	Match   []string `json:"match"`
	Next    []*Node  `json:"next,omitempty"`
}

// Read a Config object from r.
func LoadConfig(r io.Reader) (Config, error) {
	var ret []*Node
	dec := json.NewDecoder(r)
	err := dec.Decode(&ret)
	if err == nil {
		for i := range ret {
			ret[i].normalize()
		}
	}
	return ret, err
}

// Fill in defaults that were not specified in the config file.
func (n *Node) normalize() {
	if n.User == "" {
		n.User = os.Getenv("USER")
	}
	if n.Port == 0 {
		n.Port = 22
	}
	for i := range n.Next {
		n.Next[i].normalize()
	}
}
