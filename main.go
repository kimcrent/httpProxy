package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
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
func handleConnetion(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	log.Printf("Клиент подключен: %s", clientAddr)

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Клиент %s отключился: %v", clientAddr, err)
			return
		}
		message = strings.TrimSpace(message)
		log.Printf("Получено от %s: %s", clientAddr, message)

		response := processMessage(message)

		_, err = fmt.Fprintf(conn, "%s\n", response)
		if err != nil {
			log.Printf("Ошибка отправки данных клиенту %s: %v", clientAddr, err)
			return
		}

	}

}
func processMessage(msg string) string {
	return strings.ToUpper(msg)
}
