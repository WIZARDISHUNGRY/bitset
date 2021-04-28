package bitset

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	const (
		binStr  = `00000000111111110101010110101010100000001000000001111111101010101101010101000000010000000011111111010101011010101010000000100000000111111110101010110101010100000001`
		maskStr = `1111111101111111111111111111111111111100111111111011111111111111111111111111111001`
	)
	s := set{}

	if err := s.UnmarshalText([]byte(binStr)); err != nil {
		t.Fatalf("s.UnmarshalText %+v", err)
	}
	if l := s.Len(); l < len(binStr) {
		t.Fatalf("Len %d < %d", l, len(binStr))
	}
	out, err := s.MarshalText()
	if err != nil {
		t.Fatalf("out = s.MarshalText %+v", err)
	}
	fmt.Println("binStr ", binStr)
	fmt.Println("out    ", string(out))

	mask, err := UnmarshalText([]byte(maskStr))
	if err != nil {
		t.Fatalf("mask = UnmarshalText %+v", err)
	}
	maskTxt := mask.String()
	fmt.Println("maskTxt", maskTxt)

	union := s.Union(mask)
	fmt.Println("union  ", union)

	diff := s.Difference(mask)
	fmt.Println("diff   ", diff)
}
