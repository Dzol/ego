package main

import "net"
import "io"
import "log"

func main() {
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
