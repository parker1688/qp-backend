package venuevo

import "encoding/xml"

type AGZRResp struct {
	XMLName xml.Name `xml:"result"`
	Info    string   `xml:"info,attr"`
	Msg     string   `xml:"msg,attr"`
}
