package main

type numberIterator struct {
	i int
}

func (n *numberIterator) Next() int {
	n.i += 1
	return n.i
}

func IterateNumber(i int) numberIterator {
	j := numberIterator{
		i,
	}
	return j
}
