package main

import (
	"log"
	"net"
)

func listner() {
	conn, err := net.Listen("tcp", "195.96.147.217:64902")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
	defer conn.Close()

}

func main() {
	listner()

}
