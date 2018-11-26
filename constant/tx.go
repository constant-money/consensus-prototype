package constant

type Transaction struct {
	from  Address
	to    Address
	value float64
	vin   map[Serial]float64
	vout  map[Serial]float64

	// 0: normal tx
	// 1: add a producer
	// 2: remove a producer
	// 3: new leader
	// 4: new cross-shard update from another shard
	txType int

	data interface{}
}

// only needed when txType is 4
type CrossShardData struct {
	UTXOData
	CommitteeData
	sigs []Signature
}

type UTXOData struct {
	utxo              MerkleTree
	prevExternalBlock Hash
}

type CommitteeData struct {
	leader    Address
	producers []Address
}
