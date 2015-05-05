package gonads

type Functor interface {
	Map(func(interface{}) interface{}) Functor
}
