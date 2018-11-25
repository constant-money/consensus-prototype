package common

type Coordinator struct {
	committees map[int]Committee // map[shardId]committee
}

func (n Coordinator) send(to int, msg ExternalMessage) {
	// KISS, chan will be replaced by libp2p
	// KISS, will send to a few producers in shards[to], not just leader
	c := n.committees[to]
	c.leader().external <- msg
}

func NewCoordinator() (cdnt *Coordinator) {
	for i := 0; i < NUMBER_OF_SHARDS; i++ {
		c := Committee{shardId: i}
		for j := 0; j < NUMBER_OF_PRODUCERS; j++ {
			p := NewProducer(cdnt, i)
			c.addProducer(p)
			p.start()
		}
	}
	return
}
