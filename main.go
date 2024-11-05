package main

import (
	"fmt"
	client "handin4/client"
	"net"
	"time"
)

var ports = []int{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

func main() {
	for i, port := range ports {
		if checkPortAvailability(port) {
			fmt.Printf("Starting client on port %d\n", port)

			nextPort := ports[0]

			if i < len(ports)-1 {
				nextPort = ports[i+1]

			}

			client.StartClient(port, nextPort)

			return
		}
	}

	fmt.Println("No available ports to start the client.")
}

func checkPortAvailability(port int) bool {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), time.Second*3)
	if err != nil {
		return true
	}

	return false
}
