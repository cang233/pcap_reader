package main

import (
	"pcap-reader/config"
	"pcap-reader/reader"
)

func main() {
	//init config
	config.Global.Init()
	config.Global.LimitPacketsPerFlow = 10
	//input
	pcapDir := "/root/download/data/pcaps/part-pcaps"
	savingDir := "/root/download/data/pcaps/part-pcaps/parsedPcapData"
	reader.Parse(pcapDir, savingDir)
}
