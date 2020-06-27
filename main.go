package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "FQDN of the host to scan")
	port := flag.String("port", "80", "port to scan")
	addr := *host + *port
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
}