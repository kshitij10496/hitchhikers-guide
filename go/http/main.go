package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portFlag := flag.Int("port", 8080, "port to serve on")

	flag.Parse()

	port := *portFlag

	if portEnv, found := os.LookupEnv("PORT"); found {
		var err error
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			os.Exit(1)
		}
	}

	addr := fmt.Sprintf(":%d", port)
	srv := newServer()
	fmt.Printf("starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, srv); err != nil {
		os.Exit(1)
	}
}
