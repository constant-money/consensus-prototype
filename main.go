package main

import "github.com/constant-money/consensus-prototype/constant"

func main() {
	// TODO
	for i := 0; i < constant.NUMBER_OF_SHARDS; i++ {
		for j := 0; j < constant.NUMBER_OF_PRODUCERS; j++ {
			p := constant.NewProducer(i)
			go p.Start()
		}
	}
}
