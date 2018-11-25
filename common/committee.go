package common

type Committee struct {
	shardId    int
	leaderAdr  Address
	producers  map[Address]*Producer
	candidates map[Address]*Producer
}

func (c *Committee) addProducer(p *Producer) {
	c.producers[p.address] = p
}

func (c *Committee) leader() *Producer {
	return c.producers[c.leaderAdr]
}
