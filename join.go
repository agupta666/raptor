package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func join(host string, port int) {

	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	termState, err := terminal.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	defer terminal.Restore(int(os.Stdin.Fd()), termState)

	<-connectStreams(conn, os.Stdin, os.Stdout)
}
