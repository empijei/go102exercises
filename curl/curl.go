package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	wikilin := []string{
		"http://en.wikipedia.org/wiki/Linux",
		"http://it.wikipedia.org/wiki/Linux",
		"http://de.wikipedia.org/wiki/Linux",
	}
	os.Stdout.Write(findFaster(wikilin))
}

func curlToChan(dest string, out chan<- []byte) {
	buf, err := curl(dest)
	if err != nil {
		log.Println(err)
		// Let's ignore this for now
	}
	out <- buf
}

func curl(dest string) ([]byte, error) {
	resp, err := http.Get(dest)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func findFaster(dests []string) []byte {
	// TODO implement this
}
