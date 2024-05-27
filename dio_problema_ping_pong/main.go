package main

import "time"

func pingar(c chan string) {
	c <- "ping"
}

func pongar(c chan string) {
	c <- "pong"
}

func main() {
	c := make(chan string)
	for {
		go pingar(c)
		println(<-c)
		time.Sleep(500 * time.Millisecond)
		go pongar(c)
		println(<-c)
		time.Sleep(500 * time.Millisecond)
	}
}
