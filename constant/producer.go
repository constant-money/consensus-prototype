package constant

import (
	"log"
	"math/rand"
	"time"
)

type Producer struct {
	Node
	shard Shard // producer's local copy of the shard
}

func NewProducer(shardId int) (p *Producer) {
	p.inbox = make(chan Transaction, CHANNEL_SIZE)
	p.shard = Shard{id: shardId} // TODO
	return
}

func (p *Producer) Start() {
	for tx := range p.inbox {
		// new tx
		switch tx.txType {
		case 0:
		case 1:
		case 2:
		case 3:
		case 4:
			// TODO: cast tx.data to CrossShardData struct
			// TODO: process UTXOData
			// TODO: process CommitteeData
		}
	}
}

func (p *Producer) newBlock() {

	// leader proposes a new block.  KISS.
	b := Block{}
	b.hash = Hash(len(p.shard.chain))
	b.randao = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()

	// leader asks for consensus from other producers in the committee
	// TODO: implement pBFT https: //arxiv.org/pdf/1704.02397.pdf
	consensus := true

	if consensus {

		// leader adds the block to its local shard
		p.shard.chain = append(p.shard.chain, b)

		// (internal) confirm new block with other nodes in the committee
		// TODO

		// (external) send utxos to other shards
		for shardId, utxo := range b.byReceivers() {
			// TODO
			log.Println(utxo)
			p.p2p.send(p.shard.otherCommittees[shardId].Leader, Transaction{})
		}

		// (state) reconstruct local UTXO state w/ txType 0, 4
		// TODO

		// (state) reconstruct local committee state w/ txType 1, 2, 3
		// TODO

		// (state) reconstruct global committee state w/ with txType 4
		// TODO

	}
}
