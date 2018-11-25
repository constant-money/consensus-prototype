package common

const (
	NUMBER_OF_SHARDS = 256
)

type Address uint
type Serial uint
type Hash uint
type Signature string
type MerkleTree map[Serial]float64 // map[serial]value
