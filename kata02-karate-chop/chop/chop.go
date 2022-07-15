package chop

type chopper struct {
	min       int
	max       int
	chopIndex int
	victim    []int
}

func newChopper(victim []int) *chopper {
	return &chopper{
		min:       0,
		max:       len(victim) - 1,
		chopIndex: (len(victim) - 1) / 2,
		victim:    victim,
	}
}

func (i *chopper) valueAtChop() int {
	return i.victim[i.chopIndex]
}

func (i *chopper) chop(target int) bool {
	direction := target - i.valueAtChop()

	if direction == 0 {
		return false
	} else if direction > 0 {
		return i.chopUpwards(target)
	} else {
		return i.chopDownwards(target)
	}
}

func (i *chopper) chopDownwards(target int) bool {
	if i.chopIndex == i.min {
		return false
	}
	newChop := i.chopIndex - (i.chopIndex-i.min)/2
	if newChop == i.chopIndex {
		i.chopIndex = newChop - 1
		return false

	}
	i.max = i.chopIndex
	i.chopIndex = newChop
	return true
}

func (i *chopper) chopUpwards(target int) bool {
	if i.chopIndex == i.max {
		return false
	}
	newChop := (i.max-i.chopIndex)/2 + i.chopIndex
	if newChop == i.chopIndex {
		i.chopIndex = newChop + 1
		return false
	}

	i.min = i.chopIndex
	i.chopIndex = newChop
	return true
}

func Chop(target int, searchArray []int) int {
	if len(searchArray) == 0 {
		return -1
	}

	chopper := newChopper(searchArray)

	for chopper.chop(target) {}

	if chopper.valueAtChop() == target {
		return chopper.chopIndex
	} else {
		return -1
	}
}
