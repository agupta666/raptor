package main

import (
	"fmt"
	"io"
	"net"
)

// ConnHandler is a type representing TCP connection handlers
type ConnHandler func(net.Conn)

// Pipe is an abstraction for a pipe with reader and writer
type Pipe struct {
	REnd *io.PipeReader
	WEnd *io.PipeWriter
}

// NewPipe creates a new Pipe instance
func NewPipe() *Pipe {
	r, w := io.Pipe()
	return &Pipe{REnd: r, WEnd: w}
}

func connect(inPipe, outPipe *Pipe, conn net.Conn) chan bool {
	finish := make(chan bool, 1)

	go func() {
		io.Copy(conn, inPipe.REnd)
		finish <- true
	}()

	go func() {
		io.Copy(outPipe.WEnd, conn)
		finish <- true
	}()

	return finish
}

func startListener(addr string, h ConnHandler) chan bool {
	ready := make(chan bool, 1)
	go func() {
		ln, err := net.Listen("tcp", addr)

		if err != nil {
			fmt.Println("ERROR:", err)
			ready <- true
			return
		}

		ready <- true

		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			go h(conn)
		}

	}()

	return ready
}

func serve(host string, sharePort int, joinPort int) {
	// sharePipe := NewPipe()
	// joinPipe := NewPipe()

	shareAddr := fmt.Sprintf("%s:%d", host, sharePort)
	joinAddr := fmt.Sprintf("%s:%d", host, joinPort)

	<-startListener(shareAddr, func(conn net.Conn) {
		token, err := readToken(conn)

		if err != nil {
			return
		}

		session := NewSession(conn)
		saveSession(token, session)
		<-connect(session.JoinPipe, session.SharePipe, conn)
	})

	<-startListener(joinAddr, func(conn net.Conn) {
		token, err := readToken(conn)

		if err != nil {
			return
		}

		session, err := lookup(token)

		if err != nil {
			conn.Close()
			return
		}
		<-connect(session.SharePipe, session.JoinPipe, conn)
	})

}
