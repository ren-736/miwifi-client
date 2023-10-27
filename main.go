package main

import (
	"miwifi-client/client"
	"miwifi-client/cmd"
	"os"
)

func defEnv(name string, def string) string {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	return value
}

func main() {
	cmd.Client = client.MustDial(
		defEnv("MI_IP", "192.168.31.1"),
		defEnv("MI_USER", "admin"),
		defEnv("MI_PSD", ""),
	)
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
