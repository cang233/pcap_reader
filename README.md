# pcap_reader
1. 读取pcap文件然后解析，提取特征保存到相应文件里。
2. 为了项目需要，当前只解析了包大小作为特征并保存
3. 如果自定义新功能，则查看`parser/parser.go`的`Handle()`和`extract2()`，按照extract2格式写一个函数在handle里调用即可。
4. 若涉及单个数据包头操作可借鉴`flow/flow.go/ReadPacket()`.

### excute
```
func main() {
	//init config
	config.Global.Init()
	config.Global.LimitPacketsPerFlow = 10
	//input
	pcapDir := "/root/download/data/pcaps/part-pcaps"
	savingDir := "/root/download/data/pcaps/part-pcaps/parsedPcapData"
	reader.Parse(pcapDir, savingDir)
}
```