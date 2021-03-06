package rpccommon

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
)

//Start JsonRPC server
func ServerRPC(addr string, rcvr interface{}) error {
	rpc.Register(rcvr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Printf("ServerRPC started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
//New JsonRPC client
func NewClient(addr string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
