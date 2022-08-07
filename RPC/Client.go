package main
import (
	"fmt"
	"log"
	"net/rpc"
	"./rpc_objects"
)

const serverAddress = "localhost"
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":3001")
	if err != nil {
		log.Fatal("Error dialing: ", err)

	}

	args := &rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("args error:" , err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}