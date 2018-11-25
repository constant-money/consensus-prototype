package common

type Network struct {
	shards  []Shard
	workers map[Address]Worker
}

func (n Network) crossShard(to uint, msg CrossShardMsg) {
	// keep it simple, just send to the shard leader
	n.workers[n.shards[to].leader].crossShardChan <- msg
}
