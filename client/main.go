package main

import "github.com/bzyfuzy/go-chat/cmd"

func main() {
	cmd.Execute()
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"sync"
// )

// var (
// 	peers         = make(map[string]net.Conn) // Stores peer connections
// 	mutex         = sync.Mutex{}              // To handle concurrent access
// 	listeningPort = "9000"                    // Default port
// )

// // Get the local IP address of the machine
// func getLocalIP() string {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return "Unknown"
// 	}
// 	for _, addr := range addrs {
// 		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
// 			return ipNet.IP.String()
// 		}
// 	}
// 	return "Unknown"
// }

// // Get the public IP address
// func getPublicIP() string {
// 	resp, err := http.Get("https://api64.ipify.org?format=text")
// 	if err != nil {
// 		return "Unknown"
// 	}
// 	defer resp.Body.Close()

// 	ip, err := bufio.NewReader(resp.Body).ReadString('\n')
// 	if err != nil {
// 		return "Unknown"
// 	}
// 	return strings.TrimSpace(ip)
// }

// // Start a TCP server
// func startServer() {
// 	listener, err := net.Listen("tcp", "[::]:"+listeningPort) // Supports IPv4 & IPv6

// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		return
// 	}
// 	defer listener.Close()

// 	localIP := getLocalIP()
// 	publicIP := getPublicIP()
// 	fmt.Println("🔥 Server started!")
// 	fmt.Println("📡 Local Address:  " + localIP + ":" + listeningPort)
// 	fmt.Println("🌍 Public Address: " + publicIP + ":" + listeningPort)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Connection error:", err)
// 			continue
// 		}

// 		peerAddr := conn.RemoteAddr().String()
// 		mutex.Lock()
// 		peers[peerAddr] = conn
// 		mutex.Unlock()

// 		fmt.Println("[NEW PEER CONNECTED]:", peerAddr)
// 		go handlePeer(conn, peerAddr)
// 	}
// }

// // Handle peer messages
// func handlePeer(conn net.Conn, addr string) {
// 	defer func() {
// 		mutex.Lock()
// 		delete(peers, addr)
// 		mutex.Unlock()
// 		conn.Close()
// 		fmt.Println("[PEER DISCONNECTED]:", addr)
// 	}()

// 	reader := bufio.NewReader(conn)
// 	for {
// 		msg, err := reader.ReadString('\n')
// 		if err != nil {
// 			break
// 		}
// 		fmt.Print(addr + ": " + msg)
// 		broadcastMessage(addr, msg)
// 	}
// }

// // Connect to a peer manually
// func connectToPeer(address string) {
// 	conn, err := net.Dial("tcp", address)
// 	if err != nil {
// 		fmt.Println("Failed to connect:", err)
// 		return
// 	}

// 	mutex.Lock()
// 	peers[address] = conn
// 	mutex.Unlock()

// 	fmt.Println("[CONNECTED TO PEER]:", address)
// 	go handlePeer(conn, address)
// }

// // Broadcast message to all connected peers
// func broadcastMessage(sender, msg string) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	for addr, conn := range peers {
// 		if addr != sender {
// 			_, _ = conn.Write([]byte(msg))
// 		}
// 	}
// }

// // Show list of connected peers
// func showPeers() {
// 	mutex.Lock()
// 	fmt.Println("\nConnected Peers:")
// 	for addr := range peers {
// 		fmt.Println(" -", addr)
// 	}
// 	mutex.Unlock()
// }

// func main() {
// 	go startServer() // Start the server in background

// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("\nEnter command (connect [IP:PORT] / peers / msg [text] / exit): ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)

// 		if strings.HasPrefix(input, "connect ") {
// 			address := strings.TrimPrefix(input, "connect ")
// 			go connectToPeer(address)
// 		} else if input == "peers" {
// 			showPeers()
// 		} else if strings.HasPrefix(input, "msg ") {
// 			message := strings.TrimPrefix(input, "msg ") + "\n"
// 			broadcastMessage("YOU", message)
// 		} else if input == "exit" {
// 			fmt.Println("Exiting chat...")
// 			break
// 		} else {
// 			fmt.Println("Unknown command.")
// 		}
// 	}
// }
