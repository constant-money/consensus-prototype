package constant

type Block struct {
	hash               Hash
	txs                []Transaction
	sigs               []Signature
	randao             uint64
	prevInternalBlock  Hash   // sames-shard
	prevExternalBlocks []Hash // cross-shard
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
