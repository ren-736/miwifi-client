package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"html/template"
	"miwifi-client/client"
	"os"
	"strconv"
)

var gwinfoTemplate = template.Must(template.New(`gwinfoTemplate`).Parse(`LAN IPV4: {{ .LAN_IPV4 }}
LAN MAC:  {{ .LAN_MAC }}
WAN IPV4: {{ .WAN_IPV4 }}
WAN MAC:  {{ .WAN_MAC }}
`))

// Client ...
var Client *client.MiwifiClient

// RootCmd ...
var RootCmd *cobra.Command

func init() {
	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: `A MiWIFI client used to modify router configurations`,
	}
	RootCmd = rootCmd
	gwinfoCmd := &cobra.Command{
		Use:   `gwinfo`,
		Short: `show gateway information`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			info := Client.GetInformation()
			gwinfoTemplate.Execute(os.Stdout, info)
		},
	}
	rootCmd.AddCommand(gwinfoCmd)
	portMapsCmd := &cobra.Command{
		Use:   `portmaps`,
		Short: `manage port mappings`,
	}
	rootCmd.AddCommand(portMapsCmd)
	portMapsListCmd := &cobra.Command{
		Use:   `list`,
		Short: `list all port mappings`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			rules := Client.ListPortMappings()
			fmtstr := "%-5v%-16v%-12v%-12v%-20v%-12v\n"
			fmt.Printf(fmtstr, "ID", "Name", "Protocol", "OuterPort", "InnerIP", "InnerPort")
			fmt.Println("-----------------------------------------------------------------------------------")
			for i, r := range rules {
				fmt.Printf(fmtstr, i+1, r.Name, r.Protocol, r.OuterPort, r.InnerIP, r.InnerPort)
			}
		},
	}
	portMapsCmd.AddCommand(portMapsListCmd)
	portMapsCreateCmd := &cobra.Command{
		Use:     `create <name> <protocol[1：TCP 2：UDP 3：TCP&UDP]> <outer-port> <inner-ip> <inner-port>`,
		Short:   `craete a port mapping`,
		Example: `create nginx 1 8080 192.168.1.6 80`,
		Args:    cobra.ExactArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			protocol := args[1]
			outerPort, _ := strconv.Atoi(args[2])
			innerIP := args[3]
			innerPort, _ := strconv.Atoi(args[4])
			Client.CreatePortMapping(name, protocol, outerPort, innerIP, innerPort)
		},
	}
	portMapsCmd.AddCommand(portMapsCreateCmd)
	portMapsDeleteCmd := &cobra.Command{
		Use:   `delete <out_port>`,
		Short: `delete a port mapping`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			Client.DeletePortMapping(args[0])
		},
	}
	portMapsCmd.AddCommand(portMapsDeleteCmd)
	devicesCmd := &cobra.Command{
		Use:   `devices`,
		Short: `manage devices`,
	}
	rootCmd.AddCommand(devicesCmd)
	devicesListCmd := &cobra.Command{
		Use:   `list`,
		Short: `list devices`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			devices := Client.ListDevices()
			if len(devices) == 0 {
				return
			}
			fmtstr := "%-20v%-18v%-10v%-18v\n"
			fmt.Printf(fmtstr, `Name`, `IPv4`, `Type`, `MAC`)
			fmt.Println("----------------------------------------------------------------------------------------------------")
			for _, d := range devices {
				fmt.Printf(fmtstr, d.Name, d.IP[0].IP,
					d.Type, d.Mac,
				)
			}
		},
	}
	devicesCmd.AddCommand(devicesListCmd)
}
