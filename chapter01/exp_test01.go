/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/18
 * @Last Modified by: liliangshan
 * *************************************************/

package chapter01

import (
	"fmt"
	"os"
)

func ExpTest01() {
	fmt.Println("hello world! [你好世界，我爱你中国🇨🇳]")

	var s, sep string

	// go 中 i++ 是语句而不是表达式，所以像 j=i++ 就不正确 for 循环可以用break 和 return结束
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "-"
	}

	fmt.Println(s)
}
