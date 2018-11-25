package common

type Transaction struct {
	from  Address
	to    Address
	value float64
	vin   map[Serial]float64
	vout  map[Serial]float64
}
