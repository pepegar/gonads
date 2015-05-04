package gonads

type Monoid interface {
	Apply(Monoid) Monoid
	Zero() Monoid
}
