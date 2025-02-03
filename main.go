package main

import (
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", "195.96.147.217:64902")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
	defer listner.Close()
	log.Println("Сервер запущен на порту :64902")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Printf("Ошибка при принятии соединения: %v", err)
			continue
		}
		go handleConnetion(conn)
	}

}
func handleConnetion() {
	defer conn.Close()
}
