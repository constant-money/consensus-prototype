package constant

type Shard struct {

	// the shard id
	id int

	// the blockchain data
	chain []Block

	// the utxo state
	utxo MerkleTree

	// the account state
	account map[Address]Account

	// the committee state (of this shard)
	committee Committee

	// the other committees' state (from this shard's perspective)
	otherCommittees map[int]Committee
}

func shardOf(a Address) int {
	return int(a) % NUMBER_OF_SHARDS
}
