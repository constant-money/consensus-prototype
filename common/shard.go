package common

type Shard struct {
	chain      []Block
	utxo       *MerkleTree
	leader     Address
	committees []Address
	candidates []Address
	network    Network // keep it simple, for communication
}

type CrossShardMsg struct {
	prevCrossShardBlock Hash
	utxo                MerkleTree
	signatures          []Signature
}

type SameShardMsg struct {
}

func shardOf(a Address) uint {
	return uint(a) % NUMBER_OF_SHARDS
}
