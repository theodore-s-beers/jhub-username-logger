package main

import (
	"fmt"
	"log"
	"os"
)

const logFile = "/var/log/jhub_user.log"

func main() {
	// Check if the correct number of args is provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <JUPYTERHUB_USER>\n", os.Args[0])
		os.Exit(1)
	}

	// Get JUPYTERHUB_USER arg
	jhubUser := os.Args[1]

	// Open log file in append mode
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Write username to log file
	if _, err := file.WriteString(jhubUser + "\n"); err != nil {
		log.Fatalf("Failed to write to log file: %v", err)
	}

	fmt.Println("Successfully logged:", jhubUser)
}
