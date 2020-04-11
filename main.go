package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <hostname>\n", os.Args[0])
		os.Exit(0)
	}

	addr, err := net.LookupHost(os.Args[1])
	if err != nil {
		os.Exit(0)
	}

	for i := range addr {
		io.WriteString(os.Stdout, addr[i])
		io.WriteString(os.Stdout, "\n")
	}

}
