package server

import (
	"fmt"
	"net"
	"sync/atomic"
)

var activeConnections int32

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
			fmt.Printf("[ ERROR ] Failed to accept TCP connection: %v\n", err)
			continue
		}
		atomic.AddInt32(&activeConnections, 1)
		fmt.Printf("[ ~ ] Active connections: %d\n", atomic.LoadInt32(&activeConnections))
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		atomic.AddInt32(&activeConnections, -1)
		fmt.Printf("[ ~ ] Active connections: %d\n", atomic.LoadInt32(&activeConnections))
	}()

	remoteAddr := conn.RemoteAddr().String()
	ip, _, _ := net.SplitHostPort(remoteAddr)
	hostname, err := net.LookupAddr(ip)
	if err != nil {
		hostname = append(hostname, "Unknown")
	}

	fmt.Println("[ ~ ] New TCP connection established.")
	fmt.Printf("  - IP: %s\n", ip)
	fmt.Printf("  - Hostname: %v\n", hostname)

	conn.Write([]byte("You have successfully connected to the TCP server!\n"))
	conn.Write([]byte("Your IP address: " + ip + "!\n"))
}
