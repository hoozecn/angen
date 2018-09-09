package angen

import "math"

func Swap(s *[]byte, i int) {
	tmp := (*s)[i]
	(*s)[i] = (*s)[i-1]
	(*s)[i-1] = tmp
}

func send(c chan *string, quit chan bool, v string) bool {
	select {
	case c <- &v:
		return true
	case <-quit:
		return false
	}

	return false
}

func Combine(c chan *string, quit chan bool, a, b string, offset uint64) {
	defer close(c)

	if a == "" {
		if offset == 0 {
			send(c, quit, b)
		}
		return
	}

	if b == "" {
		if offset == 0 {
			send(c, quit, a)
		}

		return
	}

	combine := []byte(a + b)
	oCursor := len(a)
	count := uint64(0)

	if offset == 0 {
		v := string(combine)
		if ret := send(c, quit, v); !ret {
			return
		}
	}

	for oCursor < len(combine) {
		for cursor := oCursor; cursor >= oCursor-len(a)+1; cursor -= 1 {
			Swap(&combine, cursor)
			count += 1
			if count >= offset {
				v := string(combine)
				if ret := send(c, quit, v); !ret {
					return
				}
			}
		}

		oCursor += 1
	}
}

type Combination struct {
	seq    string
	n      int
	offset uint64
	total  uint64
}

func NewCombination(seq string, n int) (comb *Combination) {
	comb = &Combination{
		seq: seq,
		n:   n,
	}

	comb.total = uint64(math.Pow(float64(len(seq)), float64(n)))
	return comb
}

func (c *Combination) SetOffset(n uint64) {
	c.offset = n
}

func (c *Combination) Value() string {
	var s []byte
	index := c.offset
	for i := 0; i < c.n; i++ {
		s = append([]byte{c.seq[index%uint64(len(c.seq))]}, s...)
		index = index / uint64(len(c.seq))
	}

	return string(s)
}

func NewNumberComb(n int) (*Combination) {
	return NewCombination("0123456789", n)
}

func NewCharComb(n int) (*Combination) {
	return NewCombination("abcdefghijklmnopqrstuvwxyz", n)
}

type Generator struct {
	_combCount  int
	_charComb   *Combination
	_numberComb *Combination
	_composer   chan *string
	offset      uint64
	total       uint64
	maxCount    uint64
	_closer     chan bool
}

func (g *Generator) setOffset() {
	g._composer = make(chan *string)
	if g.offset >= g.total {
		close(g._composer)
		return
	}

	offset := g.offset % uint64(g._combCount)
	index := g.offset / uint64(g._combCount)

	g._numberComb.SetOffset(index % g._numberComb.total)
	g._charComb.SetOffset(index / g._numberComb.total)
	go Combine(g._composer, g._closer, g._numberComb.Value(), g._charComb.Value(), offset)
}

func NewGenerator(noOfChar, noOfNumber int, offset, maxCount uint64) (*Generator) {
	gen := &Generator{
		offset:      offset,
		_combCount:  noOfChar*noOfNumber + 1,
		_charComb:   NewCharComb(noOfChar),
		_numberComb: NewNumberComb(noOfNumber),
		maxCount:    maxCount,
		_closer:     make(chan bool, 1),
	}

	gen.total = gen._charComb.total * gen._numberComb.total * uint64(gen._combCount)

	return gen
}

func (g *Generator) Gen() (output chan *string, closer chan bool) {
	g.setOffset()
	output = make(chan *string)
	closer = make(chan bool)
	count := uint64(0)

	go func() {
		defer close(output)
		defer func() { g._closer <- true }()

		for {
		L:
			for {
				select {
				case result := <-g._composer:
					{
						if result != nil {
							g.offset += 1
							output <- result
							count += 1
							if count >= g.maxCount {
								return
							}
						} else {
							if g.offset >= g.total {
								return
							}
							g.setOffset()
							break L
						}
					}
				case <-closer:
					{
						return
					}
				}
			}
		}
	}()

	return
}
