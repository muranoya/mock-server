package main

import "net/http"

func Decode(path, body string, header http.Header) ([]byte, error) {
	return []byte("hello world"), nil
}
