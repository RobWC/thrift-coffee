package main

import (
	"log"
	"math/rand"

	co "./src/gen-go/CoffeeService"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type CoffeeOrderHandler struct {
}

func (coh CoffeeOrderHandler) Ping() error {
	log.Println("PONG")
	return nil
}

func (coh CoffeeOrderHandler) OrderCoffee(coffee *co.Coffee) (r int32, err error) {
	log.Printf("%#v", coffee)
	return rand.Int31(), nil
}

func NewCoffeeOrderHandler() CoffeeOrderHandler {
	return CoffeeOrderHandler{}
}

func main() {
	handler := NewCoffeeOrderHandler()
	processor := co.NewCoffeeOrderProcessor(handler)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	serverTransport, err := thrift.NewTServerSocket("0.0.0.0:9090")
	if err != nil {
		log.Println(err)
	}
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	server.Serve()
}
