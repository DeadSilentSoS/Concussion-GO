package main

import (
    "fmt"
    "net"
    "os"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Handle commands from the agent here
    // For simplicity, let's echo back commands as a demonstration
    buffer := make([]byte, 1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Error reading:", err)
            return
        }

        command := string(buffer[:n])
        fmt.Println("Received command:", command)

        // You can implement command execution logic here
        // For example, executeCommand(command)

        // Respond to the agent (for demonstration, we'll echo back)
        _, err = conn.Write([]byte("Received command: " + command))
        if err != nil {
            fmt.Println("Error writing:", err)
            return
        }
    }
}

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("C2 Server listening on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}
