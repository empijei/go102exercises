package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type token struct{}

func TestFindFaster(t *testing.T) {
	tmr := time.NewTimer(2 * time.Second)
	defer tmr.Stop()

	initial := make(chan token)
	synchr := make(chan token, 3)

	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "1")
		close(initial)
		synchr <- token{}
	}))
	defer ts1.Close()
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-initial
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "2 or 3")
		synchr <- token{}
	}))
	defer ts2.Close()
	ts3 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-initial
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "2 or 3")
		synchr <- token{}
	}))
	defer ts3.Close()

	urls := []string{ts1.URL, ts2.URL, ts3.URL}
	res := findFaster(urls)
	if string(res) != "1\n" {
		t.Errorf(`got: %q, want: "1"`, res)
	}

	for i := 0; i < 3; {
		select {
		case <-synchr:
			i++
		case <-tmr.C:
			t.Errorf("timeout after %d requests", i)
		}
	}
}
