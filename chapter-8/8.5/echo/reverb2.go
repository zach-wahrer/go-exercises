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

type Echo struct {
	c     net.Conn
	wg    *sync.WaitGroup
	shout string
	delay time.Duration
}

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

func handleConn(c net.Conn) {
	defer c.Close()
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	heartbeat := make(chan struct{})

	go func() {
		for input.Scan() {
			wg.Add(1)
			heartbeat <- struct{}{}
			go echo(Echo{c, &wg, input.Text(), 1 * time.Second})
		}
		if input.Err() != nil {
			log.Print(input.Err())
		}
		wg.Wait()
	}()

	ticker := time.NewTicker(1 * time.Second)
	idle := 0
	for {
		select {
		case <-heartbeat:
			idle = 0
		case <-ticker.C:
			idle++
			if idle == 10 {
				wg.Wait()
				return
			}
		}
	}
}

func echo(packet Echo) {
	defer packet.wg.Done()
	fmt.Fprintln(packet.c, "\t", strings.ToUpper(packet.shout))
	time.Sleep(packet.delay)
	fmt.Fprintln(packet.c, "\t", packet.shout)
	time.Sleep(packet.delay)
	fmt.Fprintln(packet.c, "\t", strings.ToLower(packet.shout))
}
