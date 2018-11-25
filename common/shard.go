package common

type Shard struct {
	id    int
	chain []Block
	utxo  MerkleTree
}

type ExternalMessage struct {
	prevExternalBlock Hash
	utxo              MerkleTree
	sigs              []Signature
}

type InternalMessage struct {
}

func shardOf(a Address) int {
	return int(a) % NUMBER_OF_SHARDS
}
