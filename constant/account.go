package constant

type Account struct {
	address   Address
	utxo      MerkleTree // maybe
	variables map[string]interface{}
}
