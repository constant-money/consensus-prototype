package common

type Network struct {
	shards     []Shard
	committees map[Address]Committee
}

func (n Network) send(to uint, msg InterMessage) {
	// KISS, to be replaced with libp2p
	n.committees[n.shards[to].leader].interShard <- msg
}
