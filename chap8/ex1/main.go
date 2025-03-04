// usage:
// go run chap8/clock2/clock.go -port 8001
// TZ=UTC go run chap8/clock2/clock.go -port 8002
// go run chap8/ex1/main.go Local=localhost:8001 UTC=localhost:8002
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

type Host struct {
	Name    string
	Address string
}

func ParseHost(s string) Host {
	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		log.Fatalf("invalid host: %s", s)
	}
	h := Host{Name: parts[0], Address: parts[1]}
	return h
}

//goland:noinspection GoUnhandledErrorResult
func main() {
	hosts := []Host{}
	connections := map[string]net.Conn{}

	for _, arg := range os.Args[1:] {
		h := ParseHost(arg)
		hosts = append(hosts, h)
		conn, err := net.Dial("tcp", h.Address)
		if err != nil {
			log.Fatal(err)
		}
		connections[h.Name] = conn
		defer conn.Close()
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 0, ' ', tabwriter.Debug)
	for _, h := range hosts {
		fmt.Fprintf(w, "%s\t", h.Name)
	}
	fmt.Fprintln(w)
	w.Flush()

	for {
		outputs := make(map[string]string)
		var mutex sync.Mutex

		var wg sync.WaitGroup
		for _, h := range hosts {
			wg.Add(1)
			go func(h Host) {
				defer wg.Done()
				conn := connections[h.Name]
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					return
				}
				mutex.Lock()
				defer mutex.Unlock()
				outputs[h.Name] = strings.TrimSpace(string(buf[:n]))
			}(h)
		}
		wg.Wait()

		for _, h := range hosts {
			value, ok := outputs[h.Name]
			if !ok {
				fmt.Fprintf(w, "N/A\t")
			} else {
				fmt.Fprintf(w, "%s\t", value)
			}
		}
		fmt.Fprintln(w)
		w.Flush()

		time.Sleep(1 * time.Second)
	}
}
