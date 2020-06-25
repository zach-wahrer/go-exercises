package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

type Echo struct {
	c     net.Conn
	wg    *sync.WaitGroup
	shout string
	delay time.Duration
}

func echo(packet Echo) {
	defer packet.wg.Done()
	fmt.Fprintln(packet.c, "\t", strings.ToUpper(packet.shout))
	time.Sleep(packet.delay)
	fmt.Fprintln(packet.c, "\t", packet.shout)
	time.Sleep(packet.delay)
	fmt.Fprintln(packet.c, "\t", strings.ToLower(packet.shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		packet := Echo{c, &wg, input.Text(), 1 * time.Second}
		go echo(packet)
	}
	if input.Err() != nil {
		log.Print(input.Err())
	}

	wg.Wait()
	c.Close()

}
