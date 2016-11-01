package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

func share(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	token := generateToken()
	fmt.Println("Token:", token)
	fmt.Fprintln(conn, token)

	w, h, err := terminal.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	cmd := exec.Command(os.Getenv("SHELL"))
	cmd.Env = allEnvWithSize(w, h)

	pty, err := pty.Start(cmd)

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

	<-connectPty(conn, pty, os.Stdin, os.Stdout)
}
