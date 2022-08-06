package main
import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server..")
	listener, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for{
		conn, err := listener.Accept() //accept connections from clients
		if err != nil {
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn){
	for{
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf))
	}
}

/*
1 - listen to socket connection
2 - listen to incoming request
3 - create connection
4 - read data from connection
5 - output read data
*/