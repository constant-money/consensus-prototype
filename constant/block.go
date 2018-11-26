package constant

type Block struct {

	// standard stuff
	hash              Hash
	prevInternalBlock Hash
	txs               []Transaction
	sigs              []Signature

	// for random number generation
	randao uint64

	// cross-shard references
	prevExternalBlocks []Hash

	// cross-shard data proofs map[shard_id]proof
	xsProofs map[int]CrossShardProof
}

type CrossShardProof struct {
	// uxto to be sent to another shard
	utxo MerkleTree // TODO: maybe just the root hash?

	// signatures from the committee approving the data package
	sigs []Signature
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
