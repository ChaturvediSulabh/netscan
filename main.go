package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "FQDN of the host to scan")
	portMin := flag.Int("portMin", 1, "start scanning from this port")
	portMax := flag.Int("portMax", 1024, "port to scan")

	flag.Parse()

	var wg sync.WaitGroup
	for i := *portMin; i <= *portMax; i++ {
		wg.Add(i)
		go func(i int) {
			addr := *host + ":" + strconv.Itoa(i)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Printf("Connection refused for %v on %v\n", *host, i)
				return
			}
			fmt.Printf("Connection succeeded for %v on %v\n", *host, i)
			defer conn.Close()
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
