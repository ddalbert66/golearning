package main

import (
	"fmt"
	"net"
)

var message = make(chan string)

func main() {
	//創建服務器
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	go broadcast()

	//監聽
	for {
		fmt.Println("===>監聽中")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		//建立連接
		fmt.Println("建立連接成功!")

		//啟動處理業務的go程
		go handler(conn)
	}

}

type User struct {
	id   string
	name string
	msg  chan string
}

var allUser = make(map[string]User)

//處理實際業務
func handler(conn net.Conn) {
	fmt.Println("啟動業務...")

	clientAddr := conn.RemoteAddr().String()
	newUser := User{
		id:   clientAddr,
		name: clientAddr,
		msg:  make(chan string, 10),
	}

	allUser[newUser.id] = newUser

	go writeBackData(&newUser, conn)

	newuserInfo := fmt.Sprintf("[%s:%s] =>>>>> login成功", newUser.id, newUser.name)
	message <- newuserInfo
	for {
		buf := make([]byte, 1024)

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		conn.Write([]byte(buf))
		fmt.Println("服務器接收客戶端發送過來的數據為:", string(buf[:cnt]), ", cnt:", cnt)
	}

}

func broadcast() {
	fmt.Println("廣播go程序啟動...")
	defer fmt.Println("退出廣播...")

	for {
		//取得message資料
		info := <-message
		fmt.Println("接收到訊息", info)
		//寫入每個user
		for _, user := range allUser {
			user.msg <- info
		}
	}

}

func writeBackData(user *User, conn net.Conn) {
	fmt.Printf("%s監聽自己 %s \n", user.name)
	for data := range user.msg {
		fmt.Printf("%s寫回給自己 %s \n", user.name, data)
		conn.Write([]byte(data))
	}
}
