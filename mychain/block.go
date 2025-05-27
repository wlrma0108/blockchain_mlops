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
