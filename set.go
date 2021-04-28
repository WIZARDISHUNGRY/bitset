package bitset

import (
	"encoding"
	"fmt"
)

type Set interface {
	fmt.Stringer

	Atom(int) (value Atom)
	AtomLen() int
	Bit(int) (value bool)
	Difference(Set) Set
	//Equal(Set) bool
	//Intersection(Set) Set
	Len() int
	//OnesCount() int
	Union(Set) Set
}

type set struct {
	atoms  []Atom
	length int
}
type Atom = uint64

const (
	sizeofAtom    = 8
	bitSizeofAtom = 8 * sizeofAtom
)

var (
	_ Set                      = &set{}
	_ encoding.TextUnmarshaler = &set{}
	_ encoding.TextMarshaler   = &set{}
)

func UnmarshalText(text []byte) (Set, error) {
	s := set{}
	return &s, s.UnmarshalText(text)
}

func (s *set) UnmarshalText(text []byte) error {
	hasRem := len(text)%bitSizeofAtom > 0
	size := len(text) / bitSizeofAtom
	if hasRem {
		size++
	}
	s.atoms = make([]Atom, size, size)
	s.length = len(text)
	for i, b := range text {
		addr := i / bitSizeofAtom
		offset := i % bitSizeofAtom
		v := Atom(0)
		switch b {
		case '0':
		case '1':
			v = 1
		default:
			return fmt.Errorf("invalid binary character at %d '%b'", i, b)
		}
		s.atoms[addr] = s.atoms[addr] | (v << offset)
	}
	return nil

}

func (s *set) String() string {
	str, _ := s.MarshalText()
	return string(str)
}

func (s *set) MarshalText() (text []byte, err error) {
	out := make([]byte, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		b := byte('0')
		if s.Bit(i) {
			b = '1'
		}
		out[i] = b
	}
	return out, nil
}

func (s *set) Union(t Set) Set {
	u := set{}
	sizeS := s.AtomLen()
	size := sizeS
	u.length = s.Len()
	sizeT := t.AtomLen()
	if sizeT > size {
		size = sizeT
		u.length = t.Len()
	}
	u.atoms = make([]Atom, size)
	i := 0
	for {
		if i >= size {
			break
		}
		if i < sizeS {
			u.atoms[i] |= s.Atom(i)
		}
		if i < sizeT {
			u.atoms[i] |= t.Atom(i)
		}
		i++
	}
	return &u
}

func (s *set) Difference(t Set) Set {
	u := set{}
	sizeS := s.AtomLen()
	size := sizeS
	u.length = s.Len()
	sizeT := t.AtomLen()
	if sizeT < size {
		size = sizeT
		u.length = t.Len()
	}
	u.atoms = make([]Atom, size)
	i := 0
	for {
		if i >= size {
			break
		}
		u.atoms[i] = t.Atom(i) & s.Atom(i)
		i++
	}
	return &u
}

func (s *set) Atom(i int) Atom {
	return s.atoms[i]
}

func (s *set) AtomLen() int {
	return len(s.atoms)
}
func (s *set) Len() int {
	return s.length
}

func (s *set) Bit(i int) bool {
	base := i / bitSizeofAtom
	offset := i % bitSizeofAtom
	return (s.atoms[base]&(1<<offset) != 0)
}
