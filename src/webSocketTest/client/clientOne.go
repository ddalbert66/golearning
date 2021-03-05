package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	//创建一个拨号器，也可以用默认的 websocket.DefaultDialer
	//向服务器发送连接请求，websocket 统一使用 ws://，默认端口和http一样都是80
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/go/_ajax/chat?chat=room1", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello ithome30day"))
	if err != nil {
		log.Println(err)
		return
	}
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("receive: %s\n", msg)

}
