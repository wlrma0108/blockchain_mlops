package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block 구조체: 블록체인의 최소 단위
type Block struct {
	Index     int    // 블록 높이
	Timestamp string // 생성 시각
	Data      string // 블록에 담을 데이터(여기서는 단순 문자열)
	PrevHash  string // 이전 블록의 해시
	Hash      string // 이 블록의 해시
}

// generateBlock: 이전 블록과 새로운 데이터로 새 블록을 만든다
func generateBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

// isValidChain: 슬라이스로 저장된 체인 전체 무결성을 검사
func isValidChain(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		prev := chain[i-1]
		curr := chain[i]

		// 1) 이전 블록의 해시와 현재 블록의 PrevHash 일치 여부
		if curr.PrevHash != prev.Hash {
			return false
		}
		// 2) 현재 블록의 Hash가 재계산값과 일치하는지
		if calculateHash(curr) != curr.Hash {
			return false
		}
	}
	return true
}

// calculateHash: 블록 내용을 합쳐 SHA-256 해시를 계산
func calculateHash(b Block) string {
	record := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	// 제네시스 블록 직접 만들기
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesis.Hash = calculateHash(genesis)

	fmt.Println("Genesis Block 생성:")
	fmt.Printf("%+v\n", genesis)
}
