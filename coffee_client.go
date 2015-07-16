package main

import (
	"log"
	"net"
	"os"

	"./src/gen-go/CoffeeService"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		log.Println(err)
	}
	defer transport.Close()

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := CoffeeService.NewCoffeeOrderClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	coffee := CoffeeService.NewCoffee()
	coffee.Name = "foo"

	log.Println(coffee)

	log.Println(client.Ping())
	log.Println(client.OrderCoffee(coffee))
}
