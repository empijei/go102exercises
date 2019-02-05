package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var (
	dest   = flag.String("dest", "127.0.0.1:13337", "The destination address:port")
	listen = flag.String("listen", "13336", "The port to listen for connections")
)

func main() {
	flag.Parse()

	// Listening for connections on all interfaces
	listener, err := net.Listen("tcp", ":"+*listen)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer listener.Close()

	if err = proxy(listener, os.Stdout, *dest); err != nil {
		log.Fatalf("Error during proxy operations: %v", err)
	}
}

func proxy(listener net.Listener, dump io.Writer, dest string) error {
	for {
		local, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			// TODO implement this
		}()
	}
}
