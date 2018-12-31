/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"encoding/json"
	"fmt"
)

type JsonServer struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	JsonServers []JsonServer
}

func UseJSON() {
	var serverSlice ServerSlice

	str := `{"jsonServers":[{"ServerName": "Shanghai_VPN","serverIP":"127.0.0.1"},
						{"ServerName": "Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &serverSlice)
	fmt.Println(serverSlice)
}
