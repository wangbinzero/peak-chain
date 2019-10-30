package wire

import "io"

//消息头长度基于比特币设计
//比特币为
// magic    4byte
// command  12byte
// payload  4byte
// checksum 4byte
const MessageHeaderSize = 24

//基于比特币中指令固定长度，短指令其余必须填充0
const CommandSize = 12

//消息载体最大字节数
const MaxMessagePayload = 1 << 25 //32MB

//命令，用于描述消息类型
const (
	CmdVersion      = "version"
	CmdVerAck       = "verack"
	CmdGetAddr      = "getaddr"
	CmdAddr         = "addr"
	CmdGetBlocks    = "getblocks"
	CmdInv          = "inv"
	CmdGetData      = "getdata"
	CmdNotFound     = "notfound"
	CmdBlock        = "block"
	CmdTx           = "tx"
	CmdGetHeaders   = "getheaders"
	CmdHeaders      = "headers"
	CmdPing         = "ping"
	CmdPong         = "pong"
	CmdAlert        = "alert"
	CmdMemPool      = "mempool"
	CmdFilterClear  = "filterclear"
	CmdFilterLoad   = "filterload"
	CmdMerkleBlock  = "merkleblock"
	CmdReject       = "reject"
	CmsSendHeaders  = "sendheaders"
	CmdFeeFilter    = "feefilter"
	CmdGetCFilters  = "getcfilters"
	CmdGetCFHeaders = "getcfheaders"
	CmdGetCFCheckpt = "getcfcheckpt"
	CmdCFilter      = "cfilter"
	CmdCFHeaders    = "cfheaders"
	CmdCFCheckpt    = "cfcheckpt"
)

type MessageEncoding uint32

const (
	//默认消息编码格式
	BaseEncoding MessageEncoding = 1 << iota

	//除交易以外的消息编码使用该格式
	WithnessEncoding
)

//描述 peak 的消息接口
type Message interface {

	//编码
	PKCDecode(io.Reader, uint32, MessageEncoding) error
	//解码
	PKCEncode(io.Reader, uint32, MessageEncoding) error
	//指令
	Command() string
	//最大消息载体
	MaxPayloadLength(uint32) uint32
}
