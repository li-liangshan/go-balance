/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/19
 * @Last Modified by: liliangshan
 * *************************************************/

package main

import "go-balance/times"

func main() {
	//ExpDup()
	//ExpEcho()
	//rearch.ExpTest01()
	//rearch.Plot()

	//rearch.RestApi()

	times.TestTimes()

	var channel = make(chan int)

	<-channel
}
