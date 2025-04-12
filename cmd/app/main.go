package main

import (
        "bufio"
        "fmt"
        "net"
)

func main() {
        listen, err := net.Listen("tcp", ":8080")
        if err != nil {
                panic(err)
        }
        defer listen.Close()
        fmt.Println("Server is listening on port 8080...")

        for {
                conn, err := listen.Accept()
                if err != nil {
                        fmt.Println("Error accepting connection:", err)
                        continue
                }
                go handleConnection(conn)
        }
}

func handleConnection(conn net.Conn) {
        defer conn.Close()
        fmt.Println("New client connected:", conn.RemoteAddr())

        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
                msg := scanner.Text()
                fmt.Printf("Received: %s\n", msg)
        }
}