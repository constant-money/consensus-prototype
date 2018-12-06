package constant

// Account is an abstract that is not part of the block data. It is
// created based on the transaction data and reflects the current state
// of an address.  It could be hosted in memory or a fast (k, v) database
// like leveldb

type Account struct {

	// address
	address Address

	// private utxos
	joinSplit interface{}

	// public utxos
	utxo MerkleTree

	// data
	data map[string]string

	// permission
	owners map[string]bool
	quorum int

	// triggers
	triggers map[string](map[string]([]Trigger))

	// creator
	creator Address
}

func (a *Account) updateData(key string, value string, sigs []string) {
	if a.allowed(sigs) {
		a.data[key] = value
		a.reviewTriggers(key, value)
	}
}

func (a *Account) reviewTriggers(key string, value string) {
	for _, t := range a.triggers[key][value] {
		var txData map[string]string
		for _, k := range t.keys {
			txData[k] = a.data[k]
		}
		a.newTx(t.to, t.amt, txData)
	}
}

func (a *Account) addTrigger(
	key string,
	value string,
	to Address,
	amt float64,
	keys []string,
	sigs []string) {

	if a.allowed(sigs) {
		a.triggers[key][value] = append(a.triggers[key][value],
			Trigger{to, amt, keys})
	}
}

func (a *Account) allowed(sigs []string) bool {
	pubKeys := sigs // KISS

	count := 0
	for _, p := range pubKeys {
		if a.owners[p] {
			count++
			if count == a.quorum {
				return true
			}
		}
	}

	return false
}

func (a *Account) newTx(to Address, amt float64, txData map[string]string) {

}
