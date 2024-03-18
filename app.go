package main

import "enigma-lms/config"

func main() {
	server := config.NewServer()
	server.Run()
}
