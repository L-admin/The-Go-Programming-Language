package main

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		s.words = append(s.words, 0)	// todo: ?
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if word & (1 << uint8(i)) != 0 {
				len += 1
			}
		}
	}

	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] |= 1 << bit	// todo: fix bug
}

func (s *IntSet) Clear(){
	for _, word := range s.words {
		s.words[word] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var newIntSet IntSet
	newIntSet.words = make([]uint64, len(s.words), len(s.words))
	copy(newIntSet.words, s.words)
	return &newIntSet
}


func main() {
	b := 1
	if b {  // 编译错误
		// ...
	}
}