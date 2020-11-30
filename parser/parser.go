package parser

import (
	"fmt"
	"pcap-reader/config"
	"pcap-reader/filetools"
	"pcap-reader/flow"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

//Handle 读取一个pcap解析并将结果导出
func Handle(filePath, savingDir string) {
	handle, err := pcap.OpenOffline(filePath)
	if err != nil {
		fmt.Errorf("can not open pcap file:%s", filePath)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	mapper := flow.NewMapper(config.Global.CaptureBiFlow)

	for packet := range packetSource.Packets() {
		mapper.ReadPacket(packet)
	}
	//extract
	data := mapper.Handle(extract)
	save(data, savingDir, filePath)
}

//extract features
func extract(mmap map[string]*[]*gopacket.Packet) string {
	var datas string
	for k, v := range mmap {
		line := strings.Join(strings.Split(k, "-"), ",")
		line += ","
		var lengs []string
		fmt.Printf("extract flow [%s],length=%d.\n", k, len(*v))
		for _, p := range *v {
			lengs = append(lengs, strconv.Itoa(len((*p).Data())))
		}
		line += strings.Join(lengs, ",")
		line += "\n"
		datas += line
	}

	return datas
}

func save(data string, savingPath, rawFile string) {
	//extract rawFileName
	savingFile := filetools.CheckDirSuffix(savingPath, true) + filetools.ExtractRealName(rawFile, false) + ".txt"
	filetools.WriteFile(savingFile, data)
	fmt.Println("Data saved in " + savingFile)
}
