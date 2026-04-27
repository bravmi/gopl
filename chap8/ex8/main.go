// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// !+
func handleConn(c net.Conn) {
	defer c.Close()
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	messages := make(chan string)
	done := make(chan struct{})
	defer close(done)
	go func() {
		// NOTE: ignoring potential errors from input.Err()
		for input.Scan() {
			select {
			case messages <- input.Text():
			case <-done:
				return
			}
		}
	}()
	timeout := 3 * time.Second
	timer := time.NewTimer(timeout)
	for {
		select {
		case message := <-messages:
			wg.Add(1)
			go func(m string) {
				defer wg.Done()
				echo(c, m, 1*time.Second)
			}(message)
			timer.Reset(timeout)
		case <-timer.C:
			wg.Wait()
			fmt.Fprintln(c, "timeout")
			return
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
