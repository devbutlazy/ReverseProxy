package main

import (
	"bufio"
	"fmt"
	"net"
)

func startProxy(port int) {
	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("[ ! ] Error starting proxy:", err)
		return
	}
	defer listener.Close()

	fmt.Println("[ ~ ] Port free. Starting proxy server...")
	fmt.Printf("[ ✓ ] Proxy server started on %s:%d\n\n", getLocalIP(), port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ ! ] Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("[ + ] New connection from %s\n", clientAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("[ LOG ] %s: %s\n", clientAddr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[ ! ] Connection error:", err)
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "Unknown"
}

func isPortAvailable(port int) bool {
	fmt.Printf("[ ~ ] Checking port (%v) availability...\n", port)
	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	defer listener.Close()

	return true
}

func main() {
	var port int

	fmt.Print("[ INPUT ] Enter the proxy port\n>>> ")
	fmt.Scan(&port)

	available := isPortAvailable(port)

	if !available {
		fmt.Println("[ ! ] Port is already in use.")
	}
	startProxy(port)
}
