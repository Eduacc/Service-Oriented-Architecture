package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"soa/calc/client"
	"strings"
	"time"
)

func startClient(host string, port int32) {
	client := client.New(nil)
	client.Connect(host, port)
	go client.Receive()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		client.Send(text)
		select {
		case message := <-client.Messages:
			time.Sleep(2 * time.Second)
			fmt.Println(message)
		}
	}
}

func startServer(port int32) {

}

func main() {
	flagMode := flag.String("mode", "client", "start in client or server mode")
	flagPort := flag.Int("port", 12345, "port")
	flagHost := flag.String("host", "localhost", "string")
	flag.Parse()
	port := int32(*flagPort)
	mode := strings.ToLower(*flagMode)
	if mode == "server" {
		startServer(port)
	} else if mode == "client" {
		startClient(*flagHost, port)
	}
}
