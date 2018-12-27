/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/18
 * @Last Modified by: liliangshan
 * *************************************************/

package rearch

import (
	"fmt"
	"os"
	"strings"
)

func ExpEcho() {
	var s, sep = "", ""
	// go 中 range 产生一对值：索引和这个索引所对应的值， _表示空标识符，起丢弃作用
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "-"
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], "&"))
}
