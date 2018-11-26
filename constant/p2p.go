package constant

// KISS, TODO replace w/ libp2p later
type P2P struct {
	producers map[Address]*Producer
}

func (p2p *P2P) send(dest Address, tx Transaction) {
	p2p.producers[dest].inbox <- tx
}
