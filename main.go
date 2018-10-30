package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s)

	go serve()

	<-s
}

func serve() {
	l, _ := net.Listen("tcp", ":8080")
	for {
		c, _ := l.Accept()
		go echo(c)
	}
}

func echo(c net.Conn) {
	defer c.Close()
	n, err := io.Copy(c, c)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(n)
	}
}
