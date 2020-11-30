package parser

import (
	"pcap-reader/config"
	"pcap-reader/filetools"
	"pcap-reader/flow"
	"pcap-reader/logger"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

//Handle 读取一个pcap解析并将结果导出
func Handle(filePath, savingDir string) {
	handle, err := pcap.OpenOffline(filePath)
	if err != nil {
		logger.Logger.Printf("can not open pcap file:%s \n", filePath)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	mapper := flow.NewMapper(config.Global.CaptureBiFlow)

	for packet := range packetSource.Packets() {
		mapper.ReadPacket(packet)
	}

	//extract1
	mapper.Handle(extract, savingDir, filePath)

	//TODO 添加其他extractor
	//example
	//mapper.Handle(extract2, savingDir, filePath)
}

//savingPath:解析的特征要保存的文件位置，即pcap_reader里输入的savingDir
//rawFile:当前mmap里保存的pcap文件数据，即pcap_reader里输入的pcapDir中某个当前正处理的pcap文件abs路径。
func extract2(mmap map[string]*[]*gopacket.Packet, savingPath string, rawFile string) {
	//TODO
}

//extract 解析流里包大小特征，然后保存到文件
func extract(mmap map[string]*[]*gopacket.Packet, savingPath string, rawFile string) {
	var datas string
	for k, v := range mmap {
		line := strings.Join(strings.Split(k, "-"), ",")
		line += ","
		var lengs []string
		logger.Logger.Printf("extract flow [%s],length=%d.\n", k, len(*v))
		for _, p := range *v {
			lengs = append(lengs, strconv.Itoa(len((*p).Data())))
		}
		line += strings.Join(lengs, ",")
		line += "\n"
		datas += line
	}

	//extract rawFileName
	savingFile := filetools.CheckDirSuffix(savingPath, true) + filetools.ExtractRealName(rawFile, false) + ".txt"

	save(datas, savingFile)
}

//save write data to savingFile
func save(data string, savingFile string) {
	filetools.WriteFile(savingFile, data)
	logger.Logger.Println("Data saved in " + savingFile)
}
