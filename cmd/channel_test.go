/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cmd

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	close(queue)

	// range函数遍历每个从通道接收到的数据，因为queue再发送完两个
	// 数据之后就关闭了通道，所以这里我们range函数在接收到两个数据
	// 之后就结束了。如果上面的queue通道不关闭，那么range函数就不
	// 会结束，从而在接收第三个数据的时候就阻塞了。

	for item := range queue {
		fmt.Println(item)
	}
}
