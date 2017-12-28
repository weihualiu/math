package sort

import (
	"log"
	"testing"
)

type numbers int

func (this numbers) Compare(context Contexter) bool {
	switch context.(type) {
	case numbers:
		if this > context.(numbers) {
			return true
		}
	case person:
		if int(this) > context.(person).Age {
			return true
		}
	default:
		return false
	}

	return false
}

func (this numbers) Equals(context Contexter) bool {
	if this == context.(numbers) {
		return true
	}
	return false
}

type person struct {
	Age  int
	Name string
}

func (this person) Compare(context Contexter) bool {
	switch context.(type) {
	case numbers:
		if this.Age > int(context.(numbers)) {
			return true
		}
	case person:
		if this.Age > context.(person).Age {
			return true
		}
	default:
		return false
	}

	return false
}

func (this person) Equals(context Contexter) bool {
	return true
}

func TestNew(t *testing.T) {
	arr := []numbers{1, 5, 2, 6, 3}

	f := func(instance Sorter) {
		for _, v := range arr {
			instance.Append(v)
		}
		s := instance.Sort()
		for _, v := range s {
			log.Println(v)
		}
		if s[0].(numbers) != 1 || s[len(s)-1].(numbers) != 6 {
			t.Error("sort failed!")
		}
	}

	f(New(NewMerge))
}

func TestSort(t *testing.T) {
	// 对象混合排序
	arr := []numbers{1, 5, 2, 6, 20, 3}
	persons := []person{person{12, "test1"}, person{4, "test2"}, person{10, "test3"}, person{13, "test4"}}
	var context []Contexter
	for _, v := range arr {
		context = append(context, v)
	}
	for _, v := range persons {
		context = append(context, v)
	}
	res := Sort(NewMerge, context...)
	if res[0].(numbers) != 1 || res[len(res)-1].(numbers) != 20 {
		t.Error("sort failed!")
	}

	res = append(res, person{23, "test12"}, person{9, "test13"}, person{8, "test14"}, numbers(7))
	res1 := Sort(NewMerge, res...)
	if res1[0].(numbers) != 1 || res1[len(res1)-1].(person).Age != 23 {
		t.Error("sort failed!")
	}
}
