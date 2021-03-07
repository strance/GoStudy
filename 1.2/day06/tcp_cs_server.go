package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main()  {
	server()
}

func server()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("ERROR: accept failed, err:", err)
			return
		}

		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if n == 0 {
			log.Println("服务器检测到客户端关闭，断开连接。")
			return
		}
		if err != nil {
			log.Println("ERROR: read failed, err:", err)
			continue
		}
		if string(buf[:n]) == "exit\n" {
			log.Println("服务器检测到客户端退出，断开连接。")
			return
		}
		log.Printf("INFO: server recv data: %s", buf[:n])

		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(s))
		if err != nil {
			log.Println("ERROR: write failed, err:", err)
			continue
		}
	}
}