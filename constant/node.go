package constant

type Node struct {
	address Address
	inbox   chan Transaction // KISS, TODO replaced w/ libp2p later
	p2p     *P2P             // KISS, TODO replaced w/ libp2p later
}
