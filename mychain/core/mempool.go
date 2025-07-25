package core

import "sync"


func (mp *Mempool) Add(tx *Transaction, st *State) error {
    mp.mu.Lock()
    defer mp.mu.Unlock()
    if _, ok := mp.txs[tx.ID]; ok { return errors.New("dup") }
    if err := wallet.Verify(tx); err != nil { return err }
    if !st.CanSpend(tx) { return errors.New("insufficient") }
    mp.txs[tx.ID] = tx
    return nil
}

// Mempool 는 아직 블록에 포함되지 않은 Tx 들을 보관
type Mempool struct {
	mu  sync.RWMutex
	txs map[string]*Transaction // TxID → Tx
}

func NewMempool() *Mempool {
	return &Mempool{
		txs: make(map[string]*Transaction),
	}
}

// Add : 검증(중복·서명·잔고 등) 후 등록
func (mp *Mempool) Add(tx *Transaction) bool {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if _, exists := mp.txs[tx.ID]; exists {
		return false // 이미 존재 → 실패
	}
	// TODO: 서명·잔고 검증 로직 호출
	mp.txs[tx.ID] = tx
	return true
}

// Pending : 최대 n 개 Tx 스냅샷(정렬 필요 시 여기서)
func (mp *Mempool) Pending(n int) []*Transaction {
	mp.mu.RLock()
	defer mp.mu.RUnlock()

	out := make([]*Transaction, 0, n)
	for _, tx := range mp.txs {
		out = append(out, tx)
		if len(out) == n {
			break
		}
	}
	return out
}

// Remove : 블록에 포함된 TxID 들 삭제
func (mp *Mempool) Remove(ids []string) {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	for _, id := range ids {
		delete(mp.txs, id)
	}
}

// Size : 현재 Tx 수
func (mp *Mempool) Size() int {
	mp.mu.RLock()
	defer mp.mu.RUnlock()
	return len(mp.txs)
}
