package common

type Block struct {
	hash            Hash
	prevIntraBlock  Hash
	prevInterBlocks []Hash
	txs             []Transaction
	sigs            []Signature
	randao          uint64
}

// group utxos by receivers
func (b Block) groupByReceivers() (utxos map[uint]MerkleTree) {
	for _, tx := range b.txs {
		shard := shardOf(tx.to)
		for serial, value := range tx.vout {
			utxos[shard][serial] = value
		}
	}
	return
}
