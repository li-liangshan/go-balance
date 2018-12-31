/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func UseXML() {
	file, err := os.Open("web/servers.xml")
	defer file.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Println(v)

}
