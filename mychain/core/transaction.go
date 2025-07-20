package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Transaction struct {
	ID        string `json:"id"`        // Tx 해시 (32바이트 hex)
	Timestamp int64  `json:"timestamp"` // UnixNano
	From      string `json:"from"`      // 보낸 주소(지갑 PKH)
	To        string `json:"to"`        // 받는 주소
	Amount    int64  `json:"amount"`    // 단위: Satoshi 등
}

func NewTransaction(from, to string, amt int64) *Transaction {
	tx := &Transaction{
		Timestamp: time.Now().UnixNano(),
		From:      from,
		To:        to,
		Amount:    amt,
	}
	tx.ID = tx.hash()
	return tx
}

func (tx *Transaction) hash() string {
	enc, _ := json.Marshal(tx) // Signature 없이 직렬화
	sum := sha256.Sum256(enc)
	return hex.EncodeToString(sum[:])
}
