package main
import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"time"
	"rpcobjects"
)

func main(){
	calc := new(rpcobjects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "0.0.0.0:3001")
	if e != nil {
		log.Fatal("Starting RPC-server-listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}