package p2p

import "github.com/you/mychain/core"

type Message struct {
	Type string `json:"type"`
	Tx   *core.Transaction
}
