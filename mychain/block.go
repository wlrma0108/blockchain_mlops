package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Node struct {
	Address    string   // 내 리스닝 주소, 예: "localhost:3000"
	Blockchain []Block  // 내 체인 복사본
	Peers      []string // 연결된 동료 노드 주소 목록
}

type Message struct {
    Type  string // "NEW_BLOCK" or "CHAIN"
    Block *Block
    Chain []Block
}

msg := Message{Type: "NEW_BLOCK", Block: &newBlock}
payload, _ := json.Marshal(msg)
for _, peer := range node.Peers {
    go func(addr string) {
        conn, _ := net.Dial("tcp", addr)
        defer conn.Close()
        conn.Write(payload)
        conn.Write([]byte("\n"))
    }(peer)
}

// Block 구조체: 블록체인의 최소 단위
type Block struct {
	Index     int    // 블록 높이
	Timestamp string // 생성 시각
	Data      string // 블록에 담을 데이터(여기서는 단순 문자열)
	PrevHash  string // 이전 블록의 해시
	Hash      string // 이 블록의 해시
}

// handleConnection: 새 연결에서 체인을 받거나 보냄
func (node *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 1) 내 체인을 상대에게 전송
	data, _ := json.Marshal(node.Blockchain)
	conn.Write(data)
	conn.Write([]byte("\n"))

	// 2) 상대 체인 수신
	reader := bufio.NewReader(conn)
	theirChainJSON, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		log.Println("Read error:", err)
		return
	}

	var theirChain []Block
	if err := json.Unmarshal(theirChainJSON, &theirChain); err != nil {
		log.Println("Unmarshal error:", err)
		return
	}

	// 3) 더 긴 체인으로 교체
	if len(theirChain) > len(node.Blockchain) && isValidChain(theirChain) {
		node.Blockchain = theirChain
		log.Println("체인 업데이트: 새로운 길이", len(theirChain))
	}
}

func (node *Node) StartServer() {
	ln, err := net.Listen("tcp", node.Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("P2P 서버 시작: %s\n", node.Address)

	// 무한 루프: 새 연결이 오면 handleConnection 호출
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go node.handleConnection(conn)
	}
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

func saveChain(chain []Block) error {
    file, err := os.Create("chain.json")
    if err != nil {
        return err
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    return encoder.Encode(chain)
}

func loadChain() ([]Block, error) {
    file, err := os.Open("chain.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var chain []Block
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&chain); err != nil {
        return nil, err
    }
    return chain, nil
}

func calculateHash(b Block) string {
	record := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

var Blockchain []Block

func main() {
	// 1) 제네시스 블록 생성 및 체인에 추가
	var mutex = &sync.Mutex{}

	// 블록을 추가할 때
	mutex.Lock()
	Blockchain = append(Blockchain, newBlock)
	mutex.Unlock()

	// 연결 요청 처리에서 체인을 교체할 때
	mutex.Lock()
	Blockchain = theirChain
	mutex.Unlock()

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
