/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

func TestTcpServerRPC() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}
