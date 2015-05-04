package gonads

type Semigroup interface {
	Apply(Semigroup) Semigroup
}
