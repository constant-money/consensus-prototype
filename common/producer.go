package common

import (
	"log"
	"math/rand"
	"time"
)

type Receiver struct {
	external chan ExternalMessage
	internal chan InternalMessage
	txPool   chan Transaction
}

type Sender struct {
	network *Coordinator // KISS, replaced with libp2p later
}

type Producer struct {
	Account
	Receiver
	Sender
	shard Shard // producer's local copy of the shard
}

func NewProducer(n *Coordinator, shardId int) (p *Producer) {
	p.external = make(chan ExternalMessage, CHANNEL_SIZE)
	p.internal = make(chan InternalMessage, CHANNEL_SIZE)
	p.txPool = make(chan Transaction, CHANNEL_SIZE)
	p.network = n
	p.shard = Shard{id: shardId} // TODO
	return
}

func (p *Producer) start() {
	go p.handleExternalMessage()
	go p.handleInternalMessage()
	go p.handleTransaction()
}

func (p *Producer) handleExternalMessage() {
	for msg := range p.external {
		// new cross-shard msg
		log.Println(msg.utxo)
	}
}

func (p *Producer) handleInternalMessage() {
	for msg := range p.internal {
		// new same-shard msg
		log.Println(msg)
	}
}

func (p *Producer) handleTransaction() {
	for tx := range p.txPool {
		// new tx
		log.Println(tx)
	}
}

func (p *Producer) newBlock() {

	// leader proposes a new block (refs, txs)
	b := Block{}
	b.hash = Hash(len(p.shard.chain))                                   // KISS
	b.randao = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64() // KISS

	// leader asks for consensus from other producers in the committee
	// TODO: implement pBFT https: //arxiv.org/pdf/1704.02397.pdf
	consensus := true

	if consensus {
		// leader adds block to its local shard
		p.shard.chain = append(p.shard.chain, b)

		// update intra-shard nodes so they update their local shards too

		// update inter-shard nodes so they receive their new utxos
		for to, utxo := range b.byReceivers() {
			p.network.send(to, ExternalMessage{b.hash, utxo, nil}) // KISS
		}

		// update local UTXO state
	}
}
