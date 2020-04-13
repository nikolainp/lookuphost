package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	usage()

	addr, err := getAddress(os.Args[1])
	if err != nil {
		os.Exit(0)
	}

	io.WriteString(os.Stdout, addr)
	io.WriteString(os.Stdout, "\n")
}

func usage() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <hostname>\n", os.Args[0])
		os.Exit(0)
	}
}

func getAddress(hostName string) (string, error) {

	addr, err := net.LookupHost(hostName)
	if err != nil {
		return "", err
	}

	jsonAddr, err := json.Marshal(addr)
	if err != nil {
		return "", err
	}

	return string(jsonAddr), nil
}
