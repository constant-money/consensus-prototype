package common

type Shard struct {
	chain      []Block
	utxo       *MerkleTree
	leader     Address
	committees []Address
	candidates []Address
}

type InterMessage struct {
	prevInterBlock Hash
	utxo           MerkleTree
	sigs           []Signature
}

type IntraMessage struct {
}

func shardOf(a Address) uint {
	return uint(a) % NUMBER_OF_SHARDS
}
