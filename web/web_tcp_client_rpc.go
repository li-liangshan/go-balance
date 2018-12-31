/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"fmt"
	"log"
	"net/rpc"
)

func TestTcpClientRPC() {
	client, err := rpc.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d/%d=%d remainder: %d\n", args.A, args.B, quo.Quo, quo.Rem)
}