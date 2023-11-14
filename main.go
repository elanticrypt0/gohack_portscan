package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	// max of ports to scan
	topPorts := uint(500)

	// waitGroup
	var wg sync.WaitGroup

	ConsolePrintTitle()

	// scan_v1(topPorts)
	// scan_v2(topPorts, &wg)
	scan_v3(int(topPorts), &wg)

}

func scan_v1(topPorts uint) {
	for i := uint(1); i <= topPorts; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Println(err)
		} else {
			conn.Close()
			fmt.Printf("Port %d - Open", i)
		}

	}
}

func scan_v2(topPorts uint, wg *sync.WaitGroup) {
	for i := uint(1); i <= topPorts; i++ {
		wg.Add(1)
		go func(j uint) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				log.Println(err)
			} else {
				conn.Close()
				fmt.Printf("Port %d - Open", j)
			}
		}(i)
	}
	wg.Wait()
}

func scan_v3(topPorts int, wg *sync.WaitGroup) {
	// channel for the v3
	ports := make(chan int, 100)

	// This make a buffer to the channel of 100
	for i := 1; i <= cap(ports); i++ {
		go v3_worker(ports, wg)
	}

	for i := 1; i <= topPorts; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func v3_worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println("Port ", p)
		wg.Done()
	}
}
