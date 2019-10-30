package wire

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ProtocolVersion        uint32 = 70013
	MultipleAddressVersion uint32 = 209
	NetAddressTimeVersion  uint32 = 31402
	BIP0031Version         uint32 = 60000
	BIP0035Version         uint32 = 60002
	BIP0037Version         uint32 = 70001
	RejectVersion          uint32 = 70002
	BIP0111Version         uint32 = 70011
	SendHeadersVersion     uint32 = 70012
	FeeFilterVersion       uint32 = 70013
)

type ServiceFlag uint64

const (
	SFNodeNetwork ServiceFlag = 1 << iota
	SFNodeGetUTXO
	SFNodeBloom
	SFNodeWitness
	SFNodeXthin
	SFNodeBit5
	SFNodeCF
	SFNode2X
)

var sfStrings = map[ServiceFlag]string{
	SFNodeNetwork: "SFNodeNetwork",
	SFNodeGetUTXO: "SFNodeGetUTXO",
	SFNodeBloom:   "SFNodeBloom",
	SFNodeWitness: "SFNodeWitness",
	SFNodeXthin:   "SFNodeXthin",
	SFNodeBit5:    "SFNodeBit5",
	SFNodeCF:      "SFNodeCF",
	SFNode2X:      "SFNode2X",
}

var orderedSFStrings = []ServiceFlag{
	SFNodeNetwork,
	SFNodeGetUTXO,
	SFNodeBloom,
	SFNodeWitness,
	SFNodeXthin,
	SFNodeBit5,
	SFNodeCF,
	SFNode2X,
}

func (f ServiceFlag) String() string {
	if f == 0 {
		return "0x0"
	}
	s := ""
	for _, flag := range orderedSFStrings {
		if f&flag == flag {
			s += sfStrings[flag] + "|"
			f -= flag
		}
	}

	s = strings.TrimRight(s, "|")
	if f != 0 {
		s += "|0x" + strconv.FormatUint(uint64(f), 16)
	}
	s = strings.TrimLeft(s, "|")
	return s
}

type PeakCoinNet uint32

const (
	MainNet  PeakCoinNet = 0xd9b4bef9
	TestNet  PeakCoinNet = 0xdab5bffa
	TestNet3 PeakCoinNet = 0x0709110b
	SimNet   PeakCoinNet = 0x12141c16
)

var pnStrings = map[PeakCoinNet]string{
	MainNet:  "MainNet",
	TestNet:  "TestNet",
	TestNet3: "TestNet3",
	SimNet:   "SimNet",
}

func (n PeakCoinNet) String() string {
	if s, ok := pnStrings[n]; ok {
		return s
	}
	return fmt.Sprintf("未知的 PeakCoinNet (%d)", uint32(n))
}
