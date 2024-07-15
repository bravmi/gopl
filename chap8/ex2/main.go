// usage:
// go run chap8/ex2/main.go
// telnet localhost 2121
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":2121")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("FTP server listening on port 2121")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(conn, "Error getting current directory")
		return
	}
	conn.Write([]byte("> "))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		if scanner.Err() != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		parts := strings.Fields(msg)
		if len(parts) == 0 {
			conn.Write([]byte("Missing command\n> "))
			continue
		}
		cmd, args := parts[0], parts[1:]
		if cmd == "close" {
			fmt.Fprintln(conn, "Connection closed")
			return
		}
		resp := handleCmd(cmd, args, &cwd)
		conn.Write([]byte(resp + "\n> "))
	}
}

func handleCmd(cmd string, args []string, cwd *string) string {
	switch cmd {
	case "cd":
		return changeDir(args[0], cwd)
	case "ls":
		return listDir(*cwd)
	case "get":
		if len(args) != 1 {
			return "Usage: get <file>"
		}
		return getFile(args[0], *cwd)
	case "pwd":
		return *cwd
	case "close":
		return "Connection closed"
	case "help":
		return "Commands: cd, ls, get, pwd, close, help"
	default:
		return fmt.Sprintf("Unknown command: %s", cmd)
	}
}

func changeDir(dir string, cwd *string) string {
	newPath := filepath.Join(*cwd, dir)
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		return "Directory does not exist"
	}
	*cwd = newPath
	return "Changed directory to " + newPath
}

func listDir(cwd string) string {
	entries, err := os.ReadDir(cwd)

	if err != nil {
		return "Error reading directory"
	}
	var files []string
	for _, file := range entries {
		files = append(files, file.Name())
	}
	return strings.Join(files, "\n")
}

func getFile(name, cwd string) string {
	path := filepath.Join(cwd, name)
	data, err := os.ReadFile(path)
	if err != nil {
		return "Error reading file"
	}
	return string(data)
}
