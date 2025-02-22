package server

import (
	"fmt"
	"net"
)

func StartTCPServer(port int) error {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("[ ERROR ] Failed to start TCP server on port %d: %v", port, err)
	}
	defer listener.Close()

	fmt.Printf("[ 🗸 ] TCP server started on port %d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("[ ERROR ] Failed to accept TCP connection: %v", err)
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr().String()

	ip, _, _ := net.SplitHostPort(remoteAddr)

	hostname, err := net.LookupAddr(ip)
	if err != nil {
		hostname = append(hostname, "Unknown")
	}

	fmt.Println("[ ~ ] New TCP connection established.")
	fmt.Printf("	- IP: %s\n", ip)
	fmt.Printf("	- Hostname: %v\n", hostname)
}
