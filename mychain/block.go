package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type Node struct {
	Address    string   // 내 리스닝 주소, 예: "localhost:3000"
	Blockchain []Block  // 내 체인 복사본
	Peers      []string // 연결된 동료 노드 주소 목록
}

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

var Blockchain []Block

func main() {
	// 1) 제네시스 블록 생성 및 체인에 추가
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesis.Hash = calculateHash(genesis)
	Blockchain = append(Blockchain, genesis)

	fmt.Println("Genesis Block 생성:")
	fmt.Printf("%+v\n\n", genesis)

	// 2) 새 블록 두 개 생성·추가
	block1, err := generateBlock(Blockchain[len(Blockchain)-1], "First real block")
	if err != nil {
		log.Fatal(err)
	}
	Blockchain = append(Blockchain, block1)

	block2, err := generateBlock(Blockchain[len(Blockchain)-1], "second real block")
	if err != nil {
		log.Fatal(err)
	}
	Blockchain = append(Blockchain, block2)

	// 3) 전체 체인 출력
	for _, blk := range Blockchain {
		fmt.Printf("Index:%d Data:%q Hash:%s PrevHash:%s\n",
			blk.Index, blk.Data, blk.Hash, blk.PrevHash)
	}

	// 4) 무결성 검사 결과 출력
	fmt.Println("\nChain valid?", isValidChain(Blockchain))
}
