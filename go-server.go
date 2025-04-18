package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the root directory path (default: current directory): ")
	rootDir, _ := reader.ReadString('\n')
	rootDir = strings.TrimSpace(rootDir)
	if rootDir == "" {
		rootDir, _ = os.Getwd()
	}

	fmt.Print("Enter the host (default: localhost): ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)
	if host == "" {
		host = "localhost"
	}

	fmt.Print("Enter the port (default: 3000): ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)
	if port == "" {
		port = "3000"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: http.FileServer(http.Dir(rootDir)),
	}

	// รัน server แบบ non-blocking
	go func() {
		fmt.Println("Server is running on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting server:", err)
		}
	}()

	// channel สำหรับรอ interrupt signal หรือกด q
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press 'q' then Enter to stop the server.")

	// รอกด q หรือ signal
	go func() {
		for {
			input, _ := reader.ReadString('\n')
			if strings.TrimSpace(input) == "q" {
				quit <- syscall.SIGTERM
				break
			}
		}
	}()

	<-quit
	fmt.Println("Shutting down server...")

	// shutdown ด้วย timeout 5 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	} else {
		fmt.Println("Server stopped gracefully.")
	}
}
