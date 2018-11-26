package constant

type Shard struct {
	id              int
	chain           []Block           // local block state
	utxo            MerkleTree        // local utxo state
	committee       Committee         // local committee state
	otherCommittees map[int]Committee // local view of the world
}

func shardOf(a Address) int {
	return int(a) % NUMBER_OF_SHARDS
}
