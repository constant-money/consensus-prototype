package common

type Network struct {
	shards  []Shard
	workers map[Address]Committee
}

func (n Network) send(to uint, msg InterMessage) {
	// keep it simple, just send to the shard leader
	n.workers[n.shards[to].leader].interShard <- msg
}
