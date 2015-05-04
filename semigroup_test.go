package gonads

import "testing"

type SemigroupImpl struct {
	Value int
}

func (a SemigroupImpl) Apply(b SemigroupImpl) SemigroupImpl {
	return SemigroupImpl{
		Value: a.Value * b.Value,
	}
}

func TestApply(t *testing.T) {
	a := SemigroupImpl{Value: 2}
	b := SemigroupImpl{Value: 3}
	c := SemigroupImpl{Value: 4}

	if a.Apply(b.Apply(c)) != (a.Apply(b)).Apply(c) {
		t.Error("Semigroup.Apply doesn't hold associativity")
	}
}
