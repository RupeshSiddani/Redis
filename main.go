package main

import "github.com/RupeshSiddani/redis/server"

func main() {
	server := server.NewServer(":8080")
	server.Start()
}