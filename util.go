package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func connectStreams(conn net.Conn, in io.Reader, out io.Writer) chan bool {
	finish := make(chan bool, 1)

	go func() {
		io.Copy(out, conn)
		finish <- true
	}()

	go func() {
		io.Copy(conn, in)
		finish <- true
	}()
	return finish
}

func connectPty(conn net.Conn, pty io.ReadWriter, in io.Reader, out io.Writer) chan bool {
	finish := make(chan bool, 1)
	go func() {
		io.Copy(io.MultiWriter(out, conn), pty)
		finish <- true
	}()
	go func() {
		io.Copy(pty, in)
		finish <- true
	}()
	go func() {
		io.Copy(pty, conn)
		finish <- true
	}()

	return finish
}

func allEnvWithSize(w int, h int) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("COLUMNS=%d", w))
	env = append(env, fmt.Sprintf("LINES=%d", h))
	return env
}
