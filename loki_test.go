package loki_test

import (
	"fmt"
	"testing"

	"github.com/rdumont/loki/v1"
	"github.com/stretchr/testify/assert"
)

type Item struct {
	Name     string
	Quantity int
	Done     bool
}

type ShoppingList interface {
	Add(name string, qty int) (Item, error)
}

type fakeShoppingList struct {
	MockAdd loki.Method[
		loki.Params2[string, int],
		loki.Params2[Item, error],
	]
}

func (sl *fakeShoppingList) Add(name string, qty int) (Item, error) {
	return loki.Receive(&sl.MockAdd, loki.P2(name, qty)).Spread()
}

var sampleItem = Item{
	Name:     "some name",
	Quantity: 123,
	Done:     false,
}

func TestLoki(t *testing.T) {
	t.Run("given no set up", func(t *testing.T) {
		sl := new(fakeShoppingList)

		t.Run("when I call the function", func(t *testing.T) {
			item, err := sl.Add("apples", 4)

			t.Run("then it should return zero values", func(t *testing.T) {
				assert.NoError(t, err)
				assert.Equal(t, Item{}, item)
			})
		})
	})

	t.Run("given a set up to return for any params", func(t *testing.T) {
		sl := new(fakeShoppingList)
		sl.MockAdd.OnAny().
			Return(loki.P2(sampleItem, error(nil)))

		t.Run("when I call the function", func(t *testing.T) {
			item, err := sl.Add("apples", 4)

			t.Run("then it should return the set up values", func(t *testing.T) {
				assert.NoError(t, err)
				assert.Equal(t, sampleItem, item)
			})

			t.Run("and I call it again", func(t *testing.T) {
				item, err := sl.Add("apples", 4)

				t.Run("then it should return the set up values", func(t *testing.T) {
					assert.NoError(t, err)
					assert.Equal(t, sampleItem, item)
				})
			})
		})
	})

	t.Run("given a set up to return once for any params", func(t *testing.T) {
		sl := new(fakeShoppingList)
		sl.MockAdd.OnAny().
			Once().
			Return(loki.P2(sampleItem, error(nil)))

		t.Run("when I call the function", func(t *testing.T) {
			item, err := sl.Add("apples", 4)

			t.Run("then it should return the set up values", func(t *testing.T) {
				assert.NoError(t, err)
				assert.Equal(t, sampleItem, item)
			})

			t.Run("and I call it again", func(t *testing.T) {
				item, err := sl.Add("apples", 4)

				t.Run("then it should return zero values", func(t *testing.T) {
					assert.NoError(t, err)
					assert.Equal(t, Item{}, item)
				})
			})
		})
	})

	t.Run("given a set up to return for specific params", func(t *testing.T) {
		sl := new(fakeShoppingList)
		sl.MockAdd.
			On(loki.P2("apples", 4)).
			Return(loki.P2(sampleItem, error(nil)))

		t.Run("when I call the function with different params", func(t *testing.T) {
			item, err := sl.Add("so not apples", 4)

			t.Run("then it should return zero values", func(t *testing.T) {
				assert.NoError(t, err)
				assert.Equal(t, Item{}, item)
			})
		})

		t.Run("when I call the function with correct params", func(t *testing.T) {
			item, err := sl.Add("apples", 4)

			t.Run("then it should return the set up values", func(t *testing.T) {
				assert.NoError(t, err)
				assert.Equal(t, sampleItem, item)
			})
		})
	})

	t.Run("given a strict set up", func(t *testing.T) {
		t.Run("with no expectations", func(t *testing.T) {
			ft := new(fakeTest)
			sl := new(fakeShoppingList)
			sl.MockAdd.Strict(ft)

			t.Run("when I call the function", func(t *testing.T) {
				sl.Add("apples", 4)

				t.Run("and the test finishes", func(t *testing.T) {
					ft.Finish()

					t.Run("then the test should fail", func(t *testing.T) {
						assert.Equal(t, []string{
							"Received 1 unexpected calls:\n\t- apples (string); 4 (int)\n",
						}, ft.errors)
					})
				})
			})
		})

		t.Run("with expectation to receive only once", func(t *testing.T) {
			setup := func() (*fakeTest, *fakeShoppingList) {
				ft := new(fakeTest)
				sl := new(fakeShoppingList)
				sl.MockAdd.Strict(ft).OnAny().Once()
				return ft, sl
			}

			t.Run("when I call the function twice", func(t *testing.T) {
				ft, sl := setup()

				sl.Add("apples", 4)
				sl.Add("oranges", 5)

				t.Run("and the test finishes", func(t *testing.T) {
					ft.Finish()

					t.Run("then the test should fail", func(t *testing.T) {
						assert.Equal(t, []string{
							"Received 1 unexpected calls:\n\t- oranges (string); 5 (int)\n",
						}, ft.errors)
					})
				})
			})

			t.Run("when I don't call the function", func(t *testing.T) {
				ft, _ := setup()

				t.Run("and the test finishes", func(t *testing.T) {
					ft.Finish()

					t.Run("then the test should fail", func(t *testing.T) {
						assert.Equal(t, []string{
							"Expected 1, but got 0 calls with any arguments",
						}, ft.errors)
					})
				})
			})
		})

		t.Run("with expectation to receive any times", func(t *testing.T) {
			ft := new(fakeTest)
			sl := new(fakeShoppingList)
			sl.MockAdd.Strict(ft).OnAny()

			t.Run("when I don't call the function and the test finishes", func(t *testing.T) {
				ft.Finish()

				t.Run("then the test should fail", func(t *testing.T) {
					assert.Equal(t, []string{
						"Expected some, but got 0 calls with any arguments",
					}, ft.errors)
				})
			})
		})

		t.Run("with expectation to receive specific arguments any times", func(t *testing.T) {
			ft := new(fakeTest)
			sl := new(fakeShoppingList)
			sl.MockAdd.Strict(ft).On(loki.P2("apples", 5))

			t.Run("when I don't call the function and the test finishes", func(t *testing.T) {
				ft.Finish()

				t.Run("then the test should fail", func(t *testing.T) {
					assert.Equal(t, []string{
						"Expected some, but got 0 calls with: apples (string); 5 (int)",
					}, ft.errors)
				})
			})
		})
	})
}

type fakeTest struct {
	cleanups []func()
	failed   bool
	errors   []string
}

func (t *fakeTest) Helper() {}

func (t *fakeTest) Cleanup(fn func()) {
	t.cleanups = append(t.cleanups, fn)
}

func (t *fakeTest) Errorf(format string, args ...any) {
	t.errors = append(t.errors, fmt.Sprintf(format, args...))
}

func (t *fakeTest) Finish() {
	for i := len(t.cleanups) - 1; i >= 0; i-- {
		t.cleanups[i]()
	}
}
