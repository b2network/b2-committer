package merkle

import (
	"encoding/hex"

	mt "github.com/txaty/go-merkletree"
)

type MNode struct {
	data []byte
}

func (t MNode) Serialize() ([]byte, error) {
	return t.data, nil
}

func GenerateBlocks(params []string) (blocks []mt.DataBlock) {
	for _, param := range params {
		data, _ := hex.DecodeString(param)
		block := &MNode{data}
		blocks = append(blocks, block)
	}
	return
}
