package registers

type Registers struct {
	registers map[string]int
}

func New() *Registers {
	registers := make(map[string]int, 0)
	return &Registers{registers}
}

func (r *Registers) Get(register string) int {
	// Use zero case when not-ok
	value, _ := r.registers[register]
	return value
}

func (r *Registers) Set(register string, value int) {
	r.registers[register] = value
}

func (r *Registers) MaxRegisterValue() int {
	var max int
	for _, v := range r.registers {
		if v > max {
			max = v
		}
	}

	return max
}
