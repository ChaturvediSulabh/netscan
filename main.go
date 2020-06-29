package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "FQDN of the host to scan")
	portMin := flag.Int("portMin", 1, "start scanning from this port")
	portMax := flag.Int("portMax", 1024, "port to scan")

	flag.Parse()

	for i := *portMin; i <= *portMax; i++ {
		addr := *host + ":" + strconv.Itoa(i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("Connection refused for %v on %v\n", *host, i)
			continue
		}
		fmt.Printf("Connection succeeded for %v on %v\n", *host, i)
		defer conn.Close()
	}
}
