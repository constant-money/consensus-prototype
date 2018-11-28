package constant

type Contract struct {
	id    string
	value float64
	owner Address
}

// micro_state = map[element_key]element_value
type MicroState map[string]interface{}

// macro_state = map[micro_state_key]micro_state
type MacroState map[string]MicroState
