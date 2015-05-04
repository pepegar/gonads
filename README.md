Gonads
======

[![wercker status](https://app.wercker.com/status/10cbbdf41f890155e3a79e994b36f8c0/m "wercker status")](https://app.wercker.com/project/bykey/10cbbdf41f890155e3a79e994b36f8c0)

Gonads is a category abstractions library for Go.  Obviously, this is not the
way of programming in Go, and I doubt it is production ready.

This library has been very influenced by [Cats][] and [Fantasy Land][].

Algebras:
---------

###Semigroup

Every `Semigroup` must implement a binary operation (named `Apply` in Gonads),
that satisfies associative property.

```Go
if (a.Apply(b)).Apply(c) != a.Apply(b.Apply(c)) {
	t.Error("Semigroup doesn't hold associativity")
}
```


###Monoid

A value that implements the Monoid interface should also implement Semigroup. 
Apart of that, it should have a `Zero` element.

```Go
if b.Zero().Apply(b) !== b {
	t.Error("monoid doesnt hold left identity")
}

if b.Apply(b.Zero()) !== b {
	t.Error("monoid doesnt hold right identity")
}
```

###More to come...


[Cats]: https://github.com/funcool/cats
[Fantasy Land]: https://github.com/fantasyland/fantasy-land
