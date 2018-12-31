/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package web

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []sub    `xml:"sub"`
}

type sub struct {
	SubName string `xml:"subName"`
	SubIP   string `xml:"subIP"`
}

func GenerateXML() {
	v := &Servers{Version: "1"}

	v.Svs = append(v.Svs, sub{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, sub{"Beijing_VPN", "127.0.0.1"})
	output, err := xml.MarshalIndent(v, "  ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
