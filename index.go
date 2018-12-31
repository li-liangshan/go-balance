/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/19
 * @Last Modified by: liliangshan
 * *************************************************/

package main

import "go-balance/web"

func main() {
	//ExpDup()
	//ExpEcho()
	//rearch.ExpTest01()
	//rearch.Plot()

	//rearch.RestApi()

	//times.TestTimes()

	//web.WebTest()

	TestRPC()
}

func TestRPC() {
	var channel = make(chan int)
	//go web.TestServerRPC()
	//go web.TestClientRPC()

	go web.TestTcpServerRPC()
	go web.TestTcpClientRPC()
	<-channel
}