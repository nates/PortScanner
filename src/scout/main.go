package main

import (
	"flag"
	"log"
	"strings"
	"sync"
)

var (
	portArray []string
	open      []string

	ports    string = "22,80,443"
	host     string = "127.0.0.1"
	protocol string = "tcp"
	threads  int    = 1
)

func main() {
	flag.StringVar(&ports, "ports", ports, "Ports to scan.")
	flag.StringVar(&host, "host", host, "Host to scan ports.")
	flag.StringVar(&protocol, "protocol", protocol, "Protocol to use. [TCP | UDP]")
	flag.IntVar(&threads, "threads", threads, "Amount of threads to use.")
	flag.Parse()

	// Define array of ports
	portArray = strings.Split(ports, ",")

	// Create the wait group
	var wg sync.WaitGroup

	// Start a goroutine for each thread
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for threads to stop scanning
	wg.Wait()

	log.Printf("%v open ports => %v", len(open), open)
}

func worker(id int, wg *sync.WaitGroup) {
	if len(portArray) == 0 {
		wg.Done()
		return
	}

	// Get last port from the array of ports and shift it.
	port := portArray[len(portArray)-1]
	portArray[len(portArray)-1] = ""
	portArray = portArray[:len(portArray)-1]

	// Scan port
	result := scan(strings.ToLower(protocol), host, port)
	if result != false {
		open = append(open, port)
	}

	// Log result
	log.Printf("%v:%v => %v\n", host, port, result)

	// Restart process
	worker(id, wg)
}
