package constant

type Trigger struct {
	key      string
	valueKey string
}

type Action struct {
	// then newTx(to, amt, data[keys])
	to         Address
	amount     float64
	txDataKeys []string
}

// Rules are conditional transactions.  A rule is triggered by changes
// in the data of an account.
type Rule struct {
	Trigger
	Action
}
