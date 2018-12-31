/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero. ")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestServerRPC() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}