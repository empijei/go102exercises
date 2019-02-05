package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"
)

func TestProxy(t *testing.T) {
	remote, err := net.Listen("tcp", ":13437")
	if err != nil {
		t.Fatalf("Cannot setup remote: %v", err)
	}
	defer remote.Close()

	l, err := net.Listen("tcp", ":13436")
	if err != nil {
		t.Fatalf("Cannot setup local: %v", err)
	}

	var recorder bytes.Buffer
	go proxy(l, &recorder, "localhost:13437")

	c, err := net.Dial("tcp", "localhost:13436")
	if err != nil {
		t.Fatalf("Cannot connect to proxy: %v", err)
	}
	defer c.Close()
	r, err := remote.Accept()
	if err != nil {
		t.Fatalf("Cannot get connection from proxy: %v", err)
	}
	defer r.Close()

	checkWrite := func(w io.Writer, msg string) {
		n, err := fmt.Fprint(w, msg)
		if n != len(msg) || err != nil {
			t.Errorf("got (%d, %v), want (%d, <nil>)", n, err, len(msg))
		}
	}
	checkRead := func(r io.Reader, msg string) {
		p := make([]byte, len(msg))
		_, err := io.ReadAtLeast(r, p, len(msg))
		if err != nil {
			t.Errorf("got error during read: %v", err)
		}
	}

	msgs := []string{
		"hello server\n",
		"hello client, how is it going?\n",
		"fine, thanks\n",
		"good to know, bye!\n",
	}

	ws := []io.ReadWriter{c, r}
	var sb bytes.Buffer
	for i, m := range msgs {
		fmt.Fprint(&sb, m)
		checkWrite(ws[i%2], m)
		checkRead(ws[(i+1)%2], m)
	}
	if recorder.String() != sb.String() {
		t.Errorf("recorded: got %q want %q", recorder.String(), sb.String())
	}
}
