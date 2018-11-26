package constant

type Committee struct {
	Leader     Address
	producers  []Address
	candidates []Address
}

func (c *Committee) addProducer(a Address) {
	c.producers = append(c.producers, a)
}
