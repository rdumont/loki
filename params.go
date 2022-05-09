package loki

import (
	"github.com/golang/mock/gomock"
)

func P1[A any](a A) Params1[A] {
	return Params1[A]{
		A: a,
	}
}

type Params1[A any] struct {
	A A
}

func (p Params1[A]) Matches(x Params1[A]) bool {
	return argMatches(p.A, x.A)
}

func (p Params1[A]) Spread() A {
	return p.A
}

func (p Params1[A]) Values() []interface{} {
	return []interface{}{p.A}
}

func P2[A, B any](a A, b B) Params2[A, B] {
	return Params2[A, B]{
		A: a, B: b,
	}
}

type Params2[A, B any] struct {
	A A
	B B
}

func (p Params2[A, B]) Matches(x Params2[A, B]) bool {
	return argMatches(p.A, x.A) &&
		argMatches(p.B, x.B)
}

func (p Params2[A, B]) Spread() (A, B) {
	return p.A, p.B
}

func (p Params2[A, B]) Values() []interface{} {
	return []interface{}{p.A, p.B}
}

func P3[A, B, C any](a A, b B, c C) Params3[A, B, C] {
	return Params3[A, B, C]{
		A: a, B: b, C: c,
	}
}

type Params3[A, B, C any] struct {
	A A
	B B
	C C
}

func (p Params3[A, B, C]) Matches(x Params3[A, B, C]) bool {
	return argMatches(p.A, x.A) &&
		argMatches(p.B, x.B) &&
		argMatches(p.C, x.C)
}

func (p Params3[A, B, C]) Spread() (A, B, C) {
	return p.A, p.B, p.C
}

func (p Params3[A, B, C]) Values() []interface{} {
	return []interface{}{p.A, p.B, p.C}
}

func P4[A, B, C, D any](a A, b B, c C, d D) Params4[A, B, C, D] {
	return Params4[A, B, C, D]{
		A: a, B: b, C: c, D: d,
	}
}

type Params4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

func (p Params4[A, B, C, D]) Matches(x Params4[A, B, C, D]) bool {
	return argMatches(p.A, x.A) &&
		argMatches(p.B, x.B) &&
		argMatches(p.C, x.C) &&
		argMatches(p.D, x.D)
}

func (p Params4[A, B, C, D]) Spread() (A, B, C, D) {
	return p.A, p.B, p.C, p.D
}

func (p Params4[A, B, C, D]) Values() []interface{} {
	return []interface{}{p.A, p.B, p.C, p.D}
}

func P5[A, B, C, D, E any](a A, b B, c C, d D, e E) Params5[A, B, C, D, E] {
	return Params5[A, B, C, D, E]{
		A: a, B: b, C: c, D: d, E: e,
	}
}

type Params5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func (p Params5[A, B, C, D, E]) Matches(x Params5[A, B, C, D, E]) bool {
	return argMatches(p.A, x.A) &&
		argMatches(p.B, x.B) &&
		argMatches(p.C, x.C) &&
		argMatches(p.D, x.D) &&
		argMatches(p.E, x.E)
}

func (p Params5[A, B, C, D, E]) Spread() (A, B, C, D, E) {
	return p.A, p.B, p.C, p.D, p.E
}

func (p Params5[A, B, C, D, E]) Values() []interface{} {
	return []interface{}{p.A, p.B, p.C, p.D, p.E}
}

func argMatches[T any](x, y T) bool {
	return gomock.Eq(x).Matches(y)
}
