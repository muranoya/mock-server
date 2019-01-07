package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/muranoya/mock-server/src/config"
	"github.com/muranoya/mock-server/src/handler"
)

// revision sets git-revision on compile
var revision string

func main() {
	fmt.Printf("startup %v\n", revision)

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "config file missing")
		os.Exit(1)
	}
	if err := config.Load(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	httphandler, err := handler.NewHTTPHandler(config.Config().Endpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := http.ListenAndServe(config.Config().Network.Address, httphandler); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
