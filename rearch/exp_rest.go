/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/20
 * @Last Modified by: liliangshan
 * *************************************************/

package rearch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RestApi() {
	url := "http://def-test.blacklake.tech/v1/process_routing/routing%24%23%232%40%21%21%23%5E%5E%2612"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Org-Id", "2")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "dde9c7b3-7335-4256-9527-00b5beed7b2a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
