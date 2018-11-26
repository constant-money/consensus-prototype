package constant

type Account struct {
	// account unique identifier, 32-byte address
	address Address

	// account state
	state map[string]interface{}

	// account utxo
	utxo MerkleTree

	// can be controlled either by a sig or multisig
	signatures []Signature
	sigQuorum  int
}
