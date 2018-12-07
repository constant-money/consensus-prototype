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

	// rules
	rules map[string]([]Rule)

	// creator
	creator Address
}

func (a *Account) processTx(
	from Address,
	to Address,
	value float64,
	txType int,
	txData interface{},
	sigs []string) {

	if a.allowed(sigs) {

		switch txType {

		case TX_TRANSER:
			// add and remove utxos

		case TX_UPDATE_DATA:
			a.updateData(txData.(map[string]string))

		case TX_ADD_RULE:
			a.addRule(txData.(Rule))

		}
	}

}

func (a *Account) updateData(newData map[string]string) {
	for k, v := range newData {
		a.data[k] = v
		a.reviewRules(k, v)
	}
}

func (a *Account) addRule(t Rule) {
	a.rules[t.key] = append(a.rules[t.key], t)
}

func (a *Account) reviewRules(key string, value string) {
	for _, r := range a.rules[key] {
		if value == a.data[r.valueKey] {
			var txData map[string]string
			for _, k := range r.txDataKeys {
				txData[k] = a.data[k]
			}
			a.stageNewTx(r.to, r.amount, txData)
		}
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

func (a *Account) stageNewTx(to Address, amt float64, txData map[string]string) {

}
