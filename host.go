package gssh

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func DefaultKnowHosts() (ssh.HostKeyCallback, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	fmt.Println("home : ", home)
	loc := strings.Join([]string{home, ".ssh", "known_hosts"}, "\\")
	fmt.Println("loc : ", loc)
	return KnownHosts(loc)
}

func KnownHosts(kh string) (ssh.HostKeyCallback, error) {
	return knownhosts.New(kh)
}
