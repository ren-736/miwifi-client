package client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	//os.Setenv("MI_IP", "172.16.1.149")
	//os.Setenv("MI_USER", "admin")
	//os.Setenv("MI_PSD", "mima1234")
	//for _, env := range os.Environ() {
	//	fmt.Println(env)
	//}
	client := MustDial("172.16.1.149",
		"admin",
		"mima1234")
	//devices := client.ListDevices()
	//fmtstr := "%-20v%-18v%-10v%-18v\n"
	//fmt.Printf(fmtstr, `Name`, `IPv4`, `Type`, `MAC`)
	//fmt.Println("----------------------------------------------------------------------------------------------------")
	//for _, d := range devices {
	//	fmt.Printf(fmtstr, d.Name, d.IP[0].IP,
	//		d.Type, d.Mac,
	//	)
	//}

	//information := client.GetInformation()
	//marshal, _ := json.MarshalIndent(information, "", "    ")
	//fmt.Println(string(marshal))

	rules := client.ListPortMappings()
	fmtstr := "%-5v%-16v%-12v%-12v%-18v%-12v\n"
	fmt.Printf(fmtstr, "ID", "Name", "Protocol", "OuterPort", "InnerIP", "InnerPort")
	fmt.Println("-----------------------------------------------------------------------------------")
	for i, r := range rules {
		fmt.Printf(fmtstr, i+1, r.Name, r.Protocol, r.OuterPort, r.InnerIP, r.InnerPort)
	}

	//client.CreatePortMapping("test", "2", 2222, "192.168.31.10", 1111)
}
