package block

import (
	"peak-chain/chanhash"
	"time"
)

// 区块最大有效载体长度
//  version    4字节
//  preBlock   32字节
//  merkleRoot 32字节
//  timestamp  4字节
//  bits       4字节
//  nonce      4字节
const MaxBlockHeaderPayload = 16 + (chanhash.HashSize * 2)
const blockHeaderLen = 80

type BlockHeader struct {
	//当前区块版本号
	Version int32

	//上一区块哈希
	PreBlock chanhash.Hash

	//默克尔树节点
	MerkleRoot chanhash.Hash

	//生成区块时间
	Timestamp time.Time

	//生成区块的难度值
	Bits uint32

	//生成区块的随机值
	Nonce uint32
}

// 创建区块
func NewBlockHeader(version int32, prevHash, merkleHash *chanhash.Hash, bits, nonce uint32) *BlockHeader {
	return &BlockHeader{
		Version:    version,
		PreBlock:   *prevHash,
		MerkleRoot: *merkleHash,
		Timestamp:  time.Unix(time.Now().Unix(), 0),
		Bits:       bits,
		Nonce:      nonce,
	}
}
