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

func TestClosure(t *testing.T) {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}