package common

type Block struct {
	hash               Hash
	prevInternalBlock  Hash
	prevExternalBlocks []Hash
	txs                []Transaction
	sigs               []Signature
	randao             uint64
}

// group utxos by receivers
func (b Block) byReceivers() (utxos map[int]MerkleTree) {
	for _, tx := range b.txs {
		shard := shardOf(tx.to)
		for serial, value := range tx.vout {
			utxos[shard][serial] = value
		}
	}
	return
}
