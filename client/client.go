package main

import (
	"NonRelDB/client/handler"
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var host string
var port string
var dump bool
var restore bool
var location string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Defines host ip")
	flag.StringVar(&host, "h", "127.0.0.1", "Defines host ip")
	flag.StringVar(&port, "port", "9090", "Defines host port")
	flag.StringVar(&port, "p", "9090", "Defines host port")
	flag.BoolVar(&dump, "dump", false, "Requests db dump in json format from server")
	flag.BoolVar(&restore, "restore", false, "Restores db from dumped file")
	flag.Parse()
}

// main entry point for client.
func main() {
	c, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		fmt.Println("Could not connect to this server")
		os.Exit(1)
	}

	defer c.Close()

	if dump {
		fmt.Fprintf(c, "dump\n")
		dbDump, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			fmt.Println("Could not receive database dump from server")
			os.Exit(1)
		}

		fmt.Println(dbDump)
		return
	}

	if restore {
		fmt.Fprintf(c, "restore\n")
		dbRestore, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println("Could not restore database dump to server")
			os.Exit(1)
		}

		fmt.Fprintf(c, dbRestore)
		return
	}

	handler.HandleConnection(c)
}
