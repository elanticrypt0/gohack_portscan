package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {

	ports := make(chan int, 100)
	results := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {
			// wg.Add(1)
			ports <- i
		}
	}()

	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}

// this checks if the port on the current pc is opened
func worker(ports, result chan int) {
	for p := range ports {

		address := fmt.Sprintf("0.0.0.0:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		} else {
			conn.Close()
			result <- p
		}
	}
}
