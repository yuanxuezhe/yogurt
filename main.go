package main

import (
	//"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

const (
	subj = "weather"
)

//func main1() {
//
//	StartServer(subj, "s1")
//	StartServer(subj, "s2")
//	StartServer(subj, "s3")
//	//wait for subscribe complete
//	time.Sleep(1 * time.Second)
//
//	StartClient(subj)
//
//	select {}
//}

//tcp server 服务端代码

func main() {
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr
	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")
	//循环接收客户端的连接，创建一个协程具体去处理连接
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

//具体处理连接过程方法
func tcpPipe(conn *net.TCPConn) {
	//tcp连接的地址
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()

	var packet [0x10000]byte
	//接收并返回消息
	for {
		n, err := conn.Read(packet[:])
		if err != nil || err == io.EOF {
			return
		}
		fmt.Println(conn.RemoteAddr(), "==>", conn.LocalAddr(), string(packet[:n]))

		//time.Sleep(time.Second*3)

		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"

		b := []byte(msg)

		conn.Write(b)
	}
}
