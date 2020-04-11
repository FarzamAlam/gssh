package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/farzamalam/gssh"
)

var (
	host string
	port int
	user string
	pass string
	cmd  string
)

func init() {
	flag.StringVar(&host, "host", "192.168.0.106", "Host ip address to connect to.")
	flag.IntVar(&port, "port", 22, "Host port to connect to.")
	flag.StringVar(&user, "user", "astrologer", "username of the account")
	flag.StringVar(&pass, "pass", "", "password for the account")
	flag.StringVar(&cmd, "cmd", "", "command to excute on the remote")
}

func main() {
	flag.Parse()
	fmt.Println("Host: ", host)

	if pass == "" {
		panic("No password is provided")
	}
	if user == "" {
		panic("Empty username")
	}
	if host == "" {
		panic("Empty host")
	}
	auth := gssh.Password(pass)
	client, err := gssh.New(user, host, auth)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if cmd != "" {
		out, err := client.Run(cmd)
		if err != nil {
			fmt.Println("Error while executing the cmd : ", cmd)
			log.Println(err)
			return
		}
		fmt.Println(string(out))
	}
}
