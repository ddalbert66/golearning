package main

import (
	"fmt"
	"net"
)

var conn net.Conn

func init() {
	conn, err := net.Dial("tcp", ":8848")
	defer conn.Close()
	if err != nil {
		fmt.Println("net dial err:", err)
		return
	}
}

func main() {

	fmt.Println("client/server 連接建立成功")

	messageChan := make(chan string)

	var message string
	fmt.Printf("請輸入:")
	fmt.Scanln(message)
	messageChan <- message

	go sendMessage(messageChan)

}

func sendMessage(message <-chan string) {

	msgStr := <-message
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
