package common

import (
	"log"
	"math/rand"
	"time"
)

type Worker struct {
	crossShardChan chan CrossShardMsg
	sameShardChan  chan SameShardMsg
	txChan         chan Transaction
	shardId        uint
	network        Network
}

func (w *Worker) start() {

	go func() {
		for utxo := range w.crossShardChan {
			// new cross-shard msg
			log.Println(utxo)
		}
	}()

	go func() {
		for msg := range w.sameShardChan {
			// new same-shard msg
			log.Println(msg)
		}
	}()

	go func() {
		for tx := range w.txChan {
			// new tx
			log.Println(tx)
		}
	}()
}

func (w *Worker) newBlock() {

	shard := w.network.shards[w.shardId]

	// leader proposes a new block (refs, txs)
	b := Block{}
	b.blockHash = Hash(len(shard.chain))                                // keep it simple
	b.randao = rand.New(rand.NewSource(time.Now().UnixNano())).Uint64() // keep it simple

	// leader collects signature from members, consensus if 1/2 or more
	consensus := true

	if consensus {
		// leader adds block to its local shard
		shard.chain = append(shard.chain, b)

		// update same-shard nodes so they update their local shards too

		// update cross-shard nodes so they receive new utxos
		for dest, utxo := range b.groupByReceiver() {
			shard.network.crossShard(dest, CrossShardMsg{b.blockHash, utxo})
		}

		// update local UTXO state
	}
}
