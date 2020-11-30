#Pcap_reader
1. 读取pcap文件然后解析，提取特征保存到相应文件里。
2. 为了项目需要，当前只解析了包大小作为特征并保存
3. 如果自定义新功能，则查看parser/parser.go的Handle()和extract2()，按照extract2格式写一个函数在handle里调用即可。