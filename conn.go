package gssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

const (
	UDP string = "udp"
	TCP string = "tcp"
)

func Conn(client *Client, cfg *ssh.ClientConfig) (err error) {
	if client.Port == 0 {
		client.Port = 22
	}
	if client.Protocol == "" {
		client.Protocol = TCP
	}
	loc := fmt.Sprintf("%s:%d", client.Host, client.Port)
	client.Conn, err = ssh.Dial(client.Protocol, loc, cfg)
	return
}
