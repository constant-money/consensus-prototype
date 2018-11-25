package common

import (
	"log"
	"math/rand"
	"time"
)

type Communications struct {
	interShard chan InterMessage
	intraShard chan IntraMessage
	txPool     chan Transaction
}

type Committee struct {
	Communications
	shard   Shard // node's local copy of the shard
	network Network
}

func (c *Committee) start() {
	go c.handleNewInterMessage()
	go c.handleNewIntraMessage()
	go c.handleNewTransaction()
}

func (c *Committee) handleNewInterMessage() {
	for msg := range c.interShard {
		// new cross-shard msg
		log.Println(msg.utxo)
	}
}

func (c *Committee) handleNewIntraMessage() {
	for msg := range c.intraShard {
		// new same-shard msg
		log.Println(msg)
	}
}

func (c *Committee) handleNewTransaction() {
	for tx := range c.txPool {
		// new tx
		log.Println(tx)
	}
}

func (c *Committee) newBlock() {

	// leader proposes a new block (refs, txs)
	b := Block{}
	b.hash = Hash(len(c.shard.chain))                                   // KISS
	b.randao = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64() // KISS

	// leader collects signature from members, consensus if 1/2 or more
	consensus := true // KISS

	if consensus {
		// leader adds block to its local shard
		c.shard.chain = append(c.shard.chain, b)

		// update intra-shard nodes so they update their local shards too

		// update inter-shard nodes so they receive their new utxos
		for to, utxo := range b.groupByReceivers() {
			c.shard.network.send(to, InterMessage{b.hash, utxo, nil}) // KISS
		}

		// update local UTXO state
	}
}
