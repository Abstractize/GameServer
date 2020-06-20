package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	fmt.Println("Server Starting...")
	var wg sync.WaitGroup
	wg.Add(1)
	go consoleThread()
	wg.Wait()
}

func consoleThread() {
	var line string
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		line, _ = reader.ReadString('\n')

		if len(line) == 0 {
			running = false
			break
		} else {

		}
	}
}
