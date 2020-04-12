package main

import (
	"fmt"
	"os"
	"log"
	"github.com/farzamalam/gssh"
)

func main() {
	var cmd string
	fmt.Print("Command : ")
	fmt.Scanf("%s\n",&cmd)
	// Password : admin
	auth := gssh.Password("admin")
	// New(username, host, auth)
	client, err := gssh.New("admin","192.168.0.106",auth)
	if err != nil{
		panic(err)
	}
	if cmd != ""{
		out,err:= client.Run(cmd)
		if err != nil{
			log.Println("Error while executing the cmd ->",cmd, " : ", err)
			os.Exit(1)
		}
		fmt.Println("[admin] $ ", string(out))
	}
}