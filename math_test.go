package math

import (
	"github.com/weihualiu/math/sort"
	"testing"
)

type person struct {
	Age  int
	Name string
}

func (this person) Compare(context sort.Contexter) bool {
	if this.Age > context.(person).Age {
		return true
	}
	return false
}

func (this person) Equals(context sort.Contexter) bool {
	if this.Age == context.(person).Age {
		return true
	}
	return false
}

func TestSort(t *testing.T) {
	parr := []person{person{13, "test1"}, person{2, "test2"}, person{45, "test3"},
		person{23, "test4"}}

	var context []sort.Contexter
	for _, v := range parr {
		context = append(context, v)
	}
	sorts := sort.New(sort.NewBubble)
	sorts.Append(context...)
	res := sorts.Sort()
	// res := sort.Sort(sort.NewBubble, context...)

	if res[len(res)-1].(person).Age != 45 || res[0].(person).Age != 2 {
		t.Error("sort failed")
	}
}
