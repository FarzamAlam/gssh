package gssh

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func DefaultKnowHosts() (ssh.HostKeyCallback, error) {
	home := os.Getenv("HOME")
	fmt.Println("home : ", home)
	loc := strings.Join([]string{home, ".ssh", "known_hosts"}, "\\")
	fmt.Println("loc : ", loc)
	return KnownHosts(loc)
}

func KnownHosts(kh string) (ssh.HostKeyCallback, error) {
	return knownhosts.New(kh)
}
