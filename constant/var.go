package constant

type Contract struct {
	id    string
	value float64
	owner Address
}

// state_set = map[element_key]element_value
type StateSet map[string]interface{}

// state = map[state_key](state_set)
type State map[string]StateSet
