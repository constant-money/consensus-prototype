package constant

const (
	NUMBER_OF_SHARDS    = 256
	NUMBER_OF_PRODUCERS = 100
	CHANNEL_SIZE        = 10
)

const (
	TX_TRANSER = iota
	TX_UPDATE_DATA
	TX_ADD_RULE
)

type Address int
type Serial int
type Hash int
type Signature string
type MerkleTree map[Serial]float64 // map[serial]value
