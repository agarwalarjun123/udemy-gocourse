package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct {
}

type fileWriter struct {
}

func (b logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
