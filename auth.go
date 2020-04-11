package gssh

import (
	"golang.org/x/crypto/ssh"
)

type Auth []ssh.AuthMethod

func Password(pass string) Auth {
	return Auth{
		ssh.Password(pass),
	}
}
