package gonads

import "testing"

type MonoidImpl struct {
	Value int
}

func (a MonoidImpl) Apply(b MonoidImpl) MonoidImpl {
	return MonoidImpl{
		Value: a.Value + b.Value,
	}
}

func (a MonoidImpl) Zero() MonoidImpl {
	return MonoidImpl{
		Value: 0,
	}
}

func TestZero(t *testing.T) {
	a := MonoidImpl{5}

	if a.Apply(a.Zero()) != a {
		t.Error("MonoidImpl doesn't hold right identity")
	}

	if a.Zero().Apply(a) != a {
		t.Error("MonoidImpl doesn't hold left identity")
	}
}
