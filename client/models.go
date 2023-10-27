package client

import (
	"encoding/json"
	"fmt"
	"net"
)

type LoginInfo struct {
	URL   string `json:"url"`
	TOKEN string `json:"token"`
}

// RespCode ...
type RespCode struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

// PortForward
type PortForward struct {
	List []PortMappingRule `json:"list"`
}

// PortMappingRule ...
type PortMappingRule struct {
	Name      string `json:"name"`
	Protocol  int    `json:"proto"`
	OuterPort int    `json:"srcport"`
	InnerIP   string `json:"destip"`
	InnerPort string `json:"destport"`
}

// Information ...
type Information struct {
	LAN_IPV4 net.IP
	LAN_MAC  string
	WAN_IPV4 net.IP
	WAN_MAC  string
}

type _Information struct {
	LAN _IP `json:"lan"`
	WAN _IP `json:"wan"`
}

type _IP struct {
	MAC  string `json:"mac"`
	IPV4 []struct {
		IP string `json:"ip"`
	}
}

// ToGatewayInfo ...
func (gw _Information) ToInformation() Information {
	info := Information{
		LAN_IPV4: net.ParseIP(gw.LAN.IPV4[0].IP),
		LAN_MAC:  gw.LAN.MAC,
		WAN_IPV4: net.ParseIP(gw.WAN.IPV4[0].IP),
		WAN_MAC:  gw.WAN.MAC,
	}
	return info
}

// DeviceListInfo ...
type DeviceListInfo struct {
	Devices []Device `json:"list"`
}

// Device
type Device struct {
	Mac        string `json:"mac"`
	Oname      string `json:"oname"`
	Isap       int    `json:"isap"`
	Parent     string `json:"parent"`
	Authority  `json:"authority"`
	Push       int    `json:"push"`
	Online     int    `json:"online"`
	Name       string `json:"name"`
	Times      int    `json:"times"`
	IP         []IP   `json:"ip"`
	Statistics `json:"statistics"`
	Icon       string `json:"icon"`
	Type       int    `json:"type"`
}

// Authority
type Authority struct {
	WAN     int `json:"wan"`
	Pridisk int `json:"pridisk"`
	Admin   int `json:"admin"`
	LAN     int `json:"lan"`
}

// IP
type IP struct {
	Downspeed string `json:"downspeed"`
	Online    string `json:"online"`
	Active    int    `json:"active"`
	Upspeed   string `json:"upspeed"`
	IP        string `json:"ip"`
}

// Statistics
type Statistics struct {
	Downspeed string `json:"downspeed"`
	Online    string `json:"online"`
	Upspeed   string `json:"upspeed"`
}

// UnmarshalJSON ...
func (d *Device) UnmarshalJSON(data []byte) error {
	type _D Device
	json.Unmarshal(data, (*_D)(d))
	return nil
}

func String(data interface{}) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", err)
	}
	return string(val)
}
