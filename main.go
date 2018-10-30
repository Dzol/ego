package main

import (
	"os"
	"os/signal"
	"net"
	"io"
	"log"
)

func serve() {
	x, _ := net.Listen("tcp", ":8080")
	for {
		c, _ := x.Accept()
		go echo(c)
	}
}

func echo(c net.Conn) {
	defer c.Close()
	i, err := io.Copy(c, c)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(i)
	}
}

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s)

	go serve()

	<-s
}
