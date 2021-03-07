package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	runClient()
}

func runClient() {
	// 这里指定的是服务器的IP + Port, 创建通信套接字.
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')

		// 主动写数据到服务器
		_, err = conn.Write([]byte(s))
		if err != nil {
			log.Println("ERROR: write failed, err:", err)
			return
		}

		if s == "exit\n" {
			log.Println("客户端主动退出。")
			os.Exit(1)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if n == 0 {
			log.Println("检测到服务器断开，关闭客户端。")
			return
		}
		if err != nil {
			log.Println("ERROR: read failed, err:", err)
			return
		}
		fmt.Printf("client recv data: %s", string(buf[:n]))
		time.Sleep(time.Second * 5)
	}

}
