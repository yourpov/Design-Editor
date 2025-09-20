package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// take input from path and read and display file
func ReadBanner(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Printf("%s\r\n", scanner.Text())
	}
}

// takes path, timeout and calls ReadBanner in loop
func main() {
	var (
		path    string
		timeout time.Duration
	)
	fmt.Printf("Path To File: ")
	fmt.Scanln(&path)
	fmt.Printf("Refresh Timeout (ms): ")
	fmt.Scanln(&timeout)
	go func() {
		for {
			fmt.Print("\033[H\033[2J")
			ReadBanner(path)
			time.Sleep(timeout * time.Millisecond)
		}
	}()

	select {}
}
