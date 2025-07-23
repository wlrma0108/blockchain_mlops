package main

import (
	"fmt"
	"github.com/you/mychain/core"
	"github.com/you/mychain/p2p"
)

func main() {
	mp := core.NewMempool()

	// 1) 사용자 입력·지갑 모듈 등으로 Tx 생성
	tx := core.NewTransaction("alice", "bob", 1000)

	// 2) 메모리풀 등록
	if ok := mp.Add(tx); !ok {
		fmt.Println("Tx 중복")
		return
	}
	fmt.Println("Tx added → mempool size:", mp.Size())

	// 3) 네트워크로 브로드캐스트
	p2p.BroadcastNewTx(tx)
}