package flow

import (
	"fmt"
	"pcap-reader/config"
	"strconv"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

//NewMapper a inited Mapper struct
func NewMapper(isBiFlow bool) *Mapper {
	return &Mapper{
		mmap:     make(map[string]*[]*gopacket.Packet, 0),
		isBiFlow: isBiFlow,
	}
}

//Mapper parse packet and save flow info
type Mapper struct {
	mmap map[string]*[]*gopacket.Packet
	//isBiFlow 设置是否按照流方向保存，如果是双向流，则一条流里保存双向包
	isBiFlow     bool
	countStatics int
}

//ReadPacket read packet and save it to flow it belongs to
func (m *Mapper) ReadPacket(pkt gopacket.Packet) {
	// ether layer
	etherLayer := pkt.Layer(layers.LayerTypeEthernet)
	if etherLayer == nil {
		return
	}
	// ip layer
	ipLayer := pkt.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}
	ipHead := ipLayer.(*layers.IPv4)
	// tcp/udp layer
	tcpLayer := pkt.Layer(layers.LayerTypeTCP)
	var srcPort, dstPort string
	protocol := ""
	if tcpLayer != nil {
		tcp := tcpLayer.(*layers.TCP)
		protocol = "tcp"
		srcPort = strconv.Itoa(int(tcp.SrcPort))
		dstPort = strconv.Itoa(int(tcp.DstPort))
	}
	udpLayer := pkt.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udp := tcpLayer.(*layers.UDP)
		protocol = "udp"
		srcPort = strconv.Itoa(int(udp.SrcPort))
		dstPort = strconv.Itoa(int(udp.DstPort))
	}
	//if both are nil,then protocol unknown, skip handling.
	if tcpLayer == nil && udpLayer == nil {
		return
	}

	// now handling this packet
	key := genKey(protocol, ipHead.SrcIP.String(), ipHead.DstIP.String(), srcPort, dstPort, m.isBiFlow)

	//first check flow packet size limit
	if _, ok := m.mmap[key]; ok && len(*m.mmap[key]) >= config.Global.LimitPacketsPerFlow {
		return
	}

	//save it to its flow.
	if _, ok := m.mmap[key]; !ok { // if not new
		(m.mmap[key]) = &[]*gopacket.Packet{}
	}
	*(m.mmap[key]) = append(*(m.mmap[key]), &pkt)
	//statics
	m.countStatics++
}

//Status print mapper info
func (m *Mapper) Status() {
	fmt.Println("* Mapper status:")
	fmt.Println("** handled pkts:\t", m.countStatics)
	fmt.Println("** cur mapper size:\t", len(m.mmap))
}

//Handle set Handler func and return handled string data.
func (m *Mapper) Handle(f func(map[string]*[]*gopacket.Packet) string) string {
	return f(m.mmap)
}

//genKey genenrate 5 tuple,biFlow控制是否要存单向流还是双向流,biFlow=true时需要把双向流处理为1条流
func genKey(protocol string, srcIP, dstIP string, srcPort, dstPort string, isBiFlow bool) string {
	ss := []string{protocol, srcIP, dstIP, srcPort, dstPort}
	if isBiFlow {
		//sort the ip and port
		if srcIP < dstIP {
			ss[1], ss[2] = ss[2], ss[1]
		}
		if srcPort < dstPort {
			ss[3], ss[4] = ss[4], ss[3]
		}
	}
	return strings.Join(ss, "-")
}
