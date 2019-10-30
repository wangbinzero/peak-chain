package chainhash

import (
	"encoding/hex"
	"fmt"
)

// 存储哈希的最大长度
const HashSize = 32

// 字符串哈希值为哈希长度的2倍
const MaxHashStringSize = HashSize * 2

//定义哈希长度错误类型
var ErrHashStrSize = fmt.Errorf("字符哈希值最大长度为: %v", MaxHashStringSize)

//  自定义哈希类型
type Hash [HashSize]byte

// 哈希值转换为16进制字符串
func (hs *Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hs[i], hs[HashSize-1-i] = hs[HashSize-1-i], hs[i]
	}
	return hex.EncodeToString(hs[:])
}

// 拷贝
func (hs *Hash) CloneBytes() []byte {
	nHash := make([]byte, HashSize)
	copy(nHash, hs[:])
	return nHash
}

func (hs *Hash) SetBytes(nHash []byte) error {
	nHashLen := len(nHash)
	if nHashLen != HashSize {
		return fmt.Errorf("无效的哈希长度: %v 真实哈希长度为: %v", nHashLen, HashSize)
	}
	// 将nHash 拷贝到 hash
	// copy(dst,src []byte)
	copy(hs[:], nHash)
	return nil
}

// 比较哈希内容
func (hs *Hash) IsEqual(target *Hash) bool {
	if hs == nil && target == nil {
		return true
	}
	if hs == nil || target == nil {
		return false
	}
	return *hs == *target
}

// 创建哈希值
func NewHash(nHash []byte) (*Hash, error) {
	var hash Hash
	err := hash.SetBytes(nHash)
	if err != nil {
		return nil, err
	}
	return &hash, nil
}

// 根据哈希字符串创建哈希字节码
func NewHashFromStr(hash string) (*Hash, error) {
	ret := new(Hash)
	err := Decode(ret, hash)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 将哈希值转换为16进制哈希字符串
func Decode(dst *Hash, src string) error {
	if len(src) > MaxHashStringSize {
		return ErrHashStrSize
	}

	var srcBytes []byte
	if len(src)%2 == 0 {
		srcBytes = []byte(src)
	} else {
		srcBytes = make([]byte, 1+len(src))
		srcBytes[0] = '0'
		copy(srcBytes[1:], src)
	}

	var reversedHash Hash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return err
	}

	//将临时数据拷贝到目标数据对象
	for i, b := range reversedHash[:HashSize/2] {
		dst[i], dst[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}
	return nil
}
