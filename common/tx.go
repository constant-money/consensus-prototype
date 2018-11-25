package common

type Transaction struct {
	sender   Address
	receiver Address
	amt      float64
	vin      map[Serial]float64
	vout     map[Serial]float64
}
