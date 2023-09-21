package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {

	ports_2_scan := 1024

	// this make 2 channels
	//this send the data
	ports := make(chan int, 100)
	// this is the result
	results := make(chan int)

	var openports []int

	for i := 0; i < cap(ports); i++ {
		// this makes channels
		go worker(ports, results)
	}

	// this is anom func to pass i value to the ports channel.
	go func() {
		for i := 1; i < ports_2_scan; i++ {
			// wg.Add(1)
			ports <- i
		}
	}()

	// this func gets the results.
	for i := 1; i < ports_2_scan; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	// once that the results are gattered sort the ports
	sort.Ints(openports)

	// print the results
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
