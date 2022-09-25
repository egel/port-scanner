package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sort"
)

func worker(ports, results chan int, hostname string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", hostname, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		if err := conn.Close(); err != nil {
			fmt.Println("connection could not be closed", err)
		}
		results <- p
	}
}

func main() {
	// given host
	host := os.Args[1]
	if host == "" {
		log.Fatal("no host provided")
	}

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	portLimit := 65535

	// spawning workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, host)
	}

	// assigning tasks to ports
	go func() {
		for i := 1; i <= portLimit; i++ {
			ports <- i
		}
	}()

	// reading results
	for i := 0; i < portLimit; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	// displaying results
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("port %d is open\n", port)
	}
	fmt.Printf("%d total ports open\n", len(openports))
}
