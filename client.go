package gssh

import (
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	Host     string
	Port     int
	User     string
	Auth     Auth
	Conn     *ssh.Client
	Protocol string
}

func New(user, host string, auth Auth) (*Client, error) {
	callback, err := DefaultKnowHosts()
	if err != nil {
		return nil, nil
	}
	client, err := NewConn(user, host, auth, callback)
	return client, err
}

func NewConn(user, host string, auth Auth, callback ssh.HostKeyCallback) (*Client, error) {
	client := &Client{
		User: user,
		Host: host,
		Auth: auth,
	}
	err := Conn(client, &ssh.ClientConfig{
		User:            client.User,
		Auth:            client.Auth,
		Timeout:         20 * time.Second,
		HostKeyCallback: callback,
	})
	return client, err
}

func (client Client) Close() error {
	// TOCH
	return client.Conn.Close()
}

func (client Client) Run(cmd string) ([]byte, error) {
	session, err := client.Conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session.CombinedOutput(cmd)
}
