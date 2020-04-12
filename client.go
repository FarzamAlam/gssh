package gssh

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
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

func (client Client) GetTerminal() error {
	session, err := client.Conn.NewSession()
	if err != nil {
		fmt.Errorf("Error at NewSession")
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	fileDescriptor := int(os.Stdin.Fd())
	if terminal.IsTerminal(fileDescriptor) {
		originalState, err := terminal.MakeRaw(fileDescriptor)
		if err != nil {
			fmt.Print("Error at MakeRaw")
			return err
		}
		defer terminal.Restore(fileDescriptor, originalState)
		//termWidth, termHeight, err := terminal.GetSize(fileDescriptor)
		if err != nil {
			fmt.Print("Error at GetSize:")
			return err
		}
		// Request pseudo terminal
		err = session.RequestPty("vt100", 80, 40, modes)
		if err != nil {
			fmt.Print("Error at RequestPty")
			return err
		}
	}

	err = session.Shell()
	if err != nil {
		fmt.Print("Error at Shell:")
		return err
	}
	err = session.Wait()
	if err != nil {
		fmt.Print("Error at Wait :")
	}
	return err
}
