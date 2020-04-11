package gssh

import (
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func DefaultKnowHosts() (ssh.HostKeyCallback, error) {
	return KnownHosts(strings.Join([]string{os.Getenv("HOME"), ".ssh", "known_hosts"}, "/"))
}

func KnownHosts(kh string) (ssh.HostKeyCallback, error) {
	return knownhosts.New(kh)
}
