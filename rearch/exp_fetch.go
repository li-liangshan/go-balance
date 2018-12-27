/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/19
 * @Last Modified by: liliangshan
 * *************************************************/

package rearch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		if len(url) == 0 {
			url = "http://www.baidu.com"
		}
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
