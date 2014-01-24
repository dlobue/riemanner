package main

import (
	"flag"
	"fmt"
	"github.com/Clever/riemanner/riemanner"
	"github.com/amir/raidman"
	"log"
	"os"
)

func main() {

	var useUDP bool
	var server string
	var port int
	flag.BoolVar(&useUDP, "udp", false, "Use UDP instead of the default stream connection (TCP).")
	flag.StringVar(&server, "server", "localhost", "Send events to the specified remote Riemann server. The default is localhost.")
	flag.IntVar(&port, "port", 5555, "Use the specified port. The default port is 5555")

	flag.Parse()

	proto := "tcp"
	if useUDP {
		proto = "udp"
	}
	client, err := raidman.Dial(proto, fmt.Sprintf("%s:%d", server, port))
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := riemanner.NewRiemanner(client, os.Stdin)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
