package main

import (
	"fmt"
	"net"
)

var messageChan = make(chan string, 10)

func main() {
	var conn net.Conn
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("net dial err:", err)
		return
	}
	fmt.Println("client/server 連接建立成功")

	var message string
	for {

		fmt.Printf("請輸入:")
		fmt.Scanln(&message)
		if len(message) > 0 {
			fmt.Println("message not nil")
			messageChan <- message
			go sendMessage(conn, messageChan)
		}
	}
}

func sendMessage(conn net.Conn, message <-chan string) {
	fmt.Println("sendMessage")
	var msgStr string
	msgStr = <-message

	sendData := []byte(msgStr)
	cnt, err := conn.Write(sendData)

	if err != nil {
		fmt.Println("conn.write err:", err)
		return
	}

	fmt.Println("client ===> server cnt:", cnt, ", data:", string(sendData))

	buf := make([]byte, 1024)

	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read ERR:", err)
	}

	fmt.Println("client <=== server cnt:", cnt, ", data:", string(buf[:cnt]))

}
