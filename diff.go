package bitset

import "encoding"

type DiffSet struct {
	t, u Set
}

var (
	_ Set = &DiffSet{}
	// _ encoding.TextUnmarshaler = &DiffSet{}
	_ encoding.TextMarshaler = &DiffSet{}
)

func (s *DiffSet) MarshalText() (text []byte, err error) {
	return marshalText(s)
}

func (s *DiffSet) String() string {
	str, _ := s.MarshalText()
	return string(str)
}

func (s *DiffSet) Atom(i int) Atom {
	var a Atom
	if i < s.Len() {
		a = s.t.Atom(i)
		if i < s.u.Len() {
			a &= s.u.Atom(i)
		}
	}
	return a
}

func (s *DiffSet) AtomLen() int {
	l := s.t.AtomLen()
	if tl := s.u.AtomLen(); tl < l {
		l = tl
	}
	return l
}

func (s *DiffSet) Len() int {
	l := s.t.Len()
	if tl := s.u.Len(); tl < l {
		l = tl
	}
	return l
}

func (s *DiffSet) Bit(i int) bool {
	b := s.t.Bit(i) && !s.u.Bit(i)
	return b
}
