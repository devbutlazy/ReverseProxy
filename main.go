package main

import (
	"fmt"
	"github.com/devbutlazy/ReverseProxy/server"
	"io"
	"net"
	"net/http"
)

func getPublicIp() (string, error) {
	url := "https://api.ipify.org"
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("[ ! ] Failed to get public IP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("[ ! ] Failed to get public IP, status code: %v", resp.Status)
	}

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("[ ! ] Failed to read IP response: %v", err)
	}

	return string(ip), nil
}

func isPortAvailable(port int) bool {
	fmt.Printf("[ ~ ] Checking port (%d) availability\n", port)

	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("[ ERROR ] %v\n", err)
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

	publicIp, err := getPublicIp()

	if err != nil {
		fmt.Printf("[ ! ] Failed to get public IP: %v\n", err)
	}
	fmt.Printf("[ 🗸 ] %s:%d is available\n", publicIp, port)

	server.StartTCPServer(port)
}
