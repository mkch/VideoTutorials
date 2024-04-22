package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const addr = "8888"

// StartServer starts a TCP server listening ":port".
func StartServer(ctx context.Context, port string, shutdown chan<- struct{}) {
	log.Println("Server starts")

	// 1: Use ctx as argument.
	var lc net.ListenConfig
	l, err := lc.Listen(ctx, "tcp", net.JoinHostPort("", port))
	if err != nil {
		if errors.Is(err, context.Canceled) {
			// ctx cancelled.
			log.Println("Server exits")
			close(shutdown)
			return
		}
		log.Panic(err)
	}

	// 2: Wait for ctx.Done() and close.
	go func() {
		<-ctx.Done()
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				// l is closed due to ctx cancellation.
				break
			}
			log.Printf("Server Error: %v\n", err)
		}

		// Worker goroutine for incoming client.
		go func(clientConn net.Conn) {
			defer clientConn.Close()

			// 3: Wait for ctx.Done() and set deadline.
			go func() {
				<-ctx.Done()
				clientConn.SetDeadline(time.UnixMilli(0))
			}()

			msg, err := bufio.NewReader(clientConn).ReadString('\n')
			if err != nil {
				if errors.Is(err, os.ErrDeadlineExceeded) {
					// clientConn is closed due to ctx cancellation.
					return
				}
				log.Printf("Server Error: %v\n", err)
				return
			}
			log.Printf("Server: message from %v: %v", clientConn.RemoteAddr(), msg)

			_, err = io.WriteString(clientConn, "Bye!\n")
			if err != nil {
				if errors.Is(err, os.ErrDeadlineExceeded) {
					// clientConn is closed due to ctx cancellation.
					return
				}
				log.Printf("Server Error: %v\n", err)
			}
		}(conn)
	}
	log.Println("Server exits")
	close(shutdown)
}

// Client connects to and communicate with server at "localhost:port".
func Client(port string, n int) {
	conn, err := net.Dial("tcp", net.JoinHostPort("localhost", port))
	if err != nil {
		log.Printf("Client %v Error: %v\n", n, err)
		return
	}
	defer conn.Close()

	_, err = io.WriteString(conn, "Hello!")
	if err != nil {
		log.Printf("Client %v Error: %v\n", n, err)
		return
	}

	// Simulate slow connection.
	time.Sleep(time.Second * 2)

	_, err = io.WriteString(conn, "\n")
	if err != nil {
		log.Printf("Client %v Error: %v\n", n, err)
		return
	}

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Client %v Error: %v\n", n, err)
		return
	}
	log.Printf("Client %v: message from server: %v", n, msg)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	shutdown := make(chan struct{})
	go StartServer(ctx, addr, shutdown)

	Client(addr, 1)
	Client(addr, 2)

	<-shutdown
}
