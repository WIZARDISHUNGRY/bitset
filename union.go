package bitset

import "encoding"

type UnionSet struct {
	t, u Set
}

var (
	_ Set = &UnionSet{}
	// _ encoding.TextUnmarshaler = &UnionSet{}
	_ encoding.TextMarshaler = &UnionSet{}
)

func (s *UnionSet) MarshalText() (text []byte, err error) {
	return marshalText(s)
}

func (s *UnionSet) String() string {
	str, _ := s.MarshalText()
	return string(str)
}

func (s *UnionSet) Atom(i int) Atom {
	var a Atom
	if i < s.Len() {
		a = s.t.Atom(i)
		if i < s.u.Len() {
			a |= s.u.Atom(i)
		}
	}
	return a
}

func (s *UnionSet) AtomLen() int {
	l := s.t.AtomLen()
	if tl := s.u.AtomLen(); tl > l {
		l = tl
	}
	return l
}

func (s *UnionSet) Len() int {
	l := s.t.Len()
	if tl := s.u.Len(); tl > l {
		l = tl
	}
	return l
}

func (s *UnionSet) Bit(i int) bool {
	b := s.t.Bit(i) || s.u.Bit(i)
	return b
}
