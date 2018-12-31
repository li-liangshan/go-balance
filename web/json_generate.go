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

type JsonBody struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type JsonSlice struct {
	Servers []JsonBody `json:"servers"`
}

func GenerateJSON() {
	var s JsonSlice
	s.Servers = append(s.Servers, JsonBody{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, JsonBody{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
