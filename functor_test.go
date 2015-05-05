package gonads

import "testing"

type List struct {
	Values []interface{}
}

func (l List) Map(fn func(interface{}) interface{}) List {
	var newL List = List{
		Values: []interface{}{},
	}

	for _, item := range l.Values {
		newL.Values = append(newL.Values, fn(item))
	}

	return newL
}

func TestIdentity(t *testing.T) {
	list := List{[]interface{}{1, 2, 3}}
	fn := func(i interface{}) interface{} { return i }

	if !listEqual(list.Map(fn), list) {
		t.Error("functor doesn't hold identity property")
	}
}

func TestComposition(t *testing.T) {
	list := List{[]interface{}{1, 2, 3}}
	timesTwo := func(i interface{}) interface{} {
		v := i.(int)
		return v * 2
	}
	plusOne := func(i interface{}) interface{} {
		v := i.(int)
		return v + 1
	}
	result1 := list.Map(timesTwo).Map(plusOne)
	result2 := list.Map(func(i interface{}) interface{} {
		v := i.(int)
		return plusOne(timesTwo(v))
	})

	if !listEqual(result1, result2) {
		t.Error("functor doesn't hold composition")
	}
}

func listEqual(a List, b List) bool {
	if len(a.Values) != len(b.Values) {
		return false
	}

	for i, item := range a.Values {
		if item != b.Values[i] {
			return false
		}
	}

	return true
}
