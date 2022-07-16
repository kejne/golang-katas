package chop

type chopper struct {
	victim []int
	min    int
	max    int
	chopI  int
	target int
}

func newChopper(target int, victim []int) *chopper {
	return &chopper{
		victim: victim,
		min:    0,
		max:    len(victim) - 1,
		chopI:  (len(victim) - 1) / 2,
		target: target,
	}
}

func (c *chopper) chop() int {
	foundResult, index := c.validate()
	if foundResult {
		return index
	}

	c.doChop()

	return c.chop()
}

func (c *chopper) doChop() {
	direction := c.target - c.victim[c.chopI]
	if direction > 0 {
		newChop := (c.max-c.chopI)/2 + c.chopI
		c.min = c.chopI
		c.chopI = newChop
	} else {
		newChop := (c.chopI-c.min)/2 + c.min
		c.max = c.chopI
		c.chopI = newChop
	}
}

func (c *chopper) validate() (bool, int) {
	if len(c.victim) == 0 {
		return true, -1
	}

	if c.found(c.chopI) {
		return true, c.chopI
	}
	
	// To avoid ping-pong, both are checked when one apart
	if c.max-c.min <= 1 {
		if c.found(c.min) {
			return true, c.min
		}
		if c.found(c.max) {
			return true, c.max
		}
		return true, -1
	}
	return false, 0
}

func (c* chopper) found(index int) bool {
	return c.victim[index] == c.target
}

func Chop(target int, victim []int) int {
	chopper := newChopper(target, victim)

	return chopper.chop()
}
