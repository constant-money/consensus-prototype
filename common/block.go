package common

type Block struct {
	blockHash            Hash
	prevSameShardBlock   Hash
	prevCrossShardBlocks []Hash
	transactions         []Transaction
	signatures           []Signature
	randao               uint64
}

func (b Block) groupByReceiver() (utxos map[uint]MerkleTree) {
	for _, tx := range b.transactions {
		shard := shardOf(tx.receiver)
		for serial, value := range tx.vout {
			utxos[shard][serial] = value
		}
	}
	return
}
