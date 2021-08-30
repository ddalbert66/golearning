package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	serverIp := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", serverIp, port)

	serverIp2 := "127.0.0.1"
	port2 := 8850
	address2 := fmt.Sprintf("%s:%d", serverIp2, port2)

	listner2, err := net.Listen("tcp", address2)
	listner, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.listen err:", err)
	}
	fmt.Println("監聽中...")
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("net.listen Accept err:", err)
		}

		go funcHandler(conn)
	}

	for {
		conn2, err := listner2.Accept()
		if err != nil {
			fmt.Println("net.listen Accept err:", err)
		}

		go funcHandler(conn2)
	}

}

func funcHandler(conn net.Conn) {

	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err == io.EOF {

	} else if err != nil {
		fmt.Println("conn.Read err:", err)
	}
	fmt.Println("讀取長度:%d,讀取字串:%s", len, string(buf))

	value := string(buf)
	value = strings.ToUpper(value)
	conn.Write([]byte(value))

	conn.Close()
}
