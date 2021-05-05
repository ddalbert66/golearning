package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	serverIp := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", serverIp, port)

	listner, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.listen err:", err)
	}
	fmt.Println("監聽中...")

	conn, err := listner.Accept()
	if err != nil {
		fmt.Println("net.listen Accept err:", err)
	}
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
	}
	fmt.Println("讀取長度:%d,讀取字串:%s", len, string(buf))

	value := string(buf)
	value = strings.ToUpper(value)
	conn.Write([]byte(value))

	conn.Close()

}
