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

func TestClientRPC() {
	client, err := rpc.DialHTTP("tcp", "localhost:9090")
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

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: ", err)
	}

	fmt.Printf("Arith: %d/%d=%d reminder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
