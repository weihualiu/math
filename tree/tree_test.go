package tree

import (
	"math/rand"
	"strconv"
	"testing"
)

type Person struct {
	Age  int
	Name string
}

func (this Person) Compare(context Context) bool {
	p, ok := context.(Person)
	if !ok {
		panic("not match interface!")
	}
	if this.Age > p.Age {
		return true
	}
	return false
}

func (this Person) Equals(context Context) bool {
	p, ok := context.(Person)
	if !ok {
		panic("not match interface!")
	}

	if this.Age == p.Age {
		return true
	}
	return false
}

func TestTreeDeep(t *testing.T) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 11, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 1, Name: "ZhaoLiu"})
	TreeAdd(tree, Person{Age: 3, Name: "SongQi"})
	TreeAdd(tree, Person{Age: 2, Name: "SongQi"})

	if TreeDeep(tree) != 4 {
		t.Errorf("tree deep error! deep value:", TreeDeep(tree))
	}

	tree1 := TreeRootNew()
	TreeAdd(tree1, Person{Age: 10, Name: "ZhangSan"})
	if TreeDeep(tree1) != 1 {
		t.Errorf("tree deep error! deep value:", TreeDeep(tree1))
	}

	tree2 := TreeRootNew()
	TreeAdd(tree2, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree2, Person{Age: 11, Name: "ZhangSan"})
	if TreeDeep(tree2) != 2 {
		t.Errorf("tree deep error! deep value:", TreeDeep(tree2))
	}
}

func TestTreeFind(t *testing.T) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 11, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 1, Name: "ZhaoLiu"})
	TreeAdd(tree, Person{Age: 3, Name: "SongQi"})
	TreeAdd(tree, Person{Age: 2, Name: "LiShanShan"})

	p, ok := TreeFind(tree, Person{Age: 3}).(Person)
	if !ok {
		t.Error("not find person")
	}
	if p.Name != "SongQi" {
		t.Errorf("not found person. ", p)
	}
	_, ok = TreeFind(tree, Person{Age: 23}).(Person)
	if ok {
		t.Error("find person error")
	}

}

func TestTreeDelete(t *testing.T) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})
	TreeAdd(tree, Person{Age: 17, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 25, Name: "WangWu2"})

	TreeDelete(tree, Person{Age: 16})

	p, _ := tree.Right.Content.(Person)
	if p.Age != 17 || tree.Left.Content.(Person).Age != 6 || tree.Right.Right.Content.(Person).Age != 18 ||
		tree.Right.Left.Content.(Person).Age != 12 || tree.Right.Left.Left.Content.(Person).Age != 11 {
		t.Errorf("delete error. ", p)
	}

	TreeDelete(tree, Person{Age: 12})
	if tree.Right.Content.(Person).Age != 17 ||
		tree.Right.Left.Content.(Person).Age != 11 ||
		tree.Right.Right.Right.Content.(Person).Age != 25 {
		t.Errorf("delete2 error. %p", p)
	}

}

func TestTreeFindMax(t *testing.T) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})

	p, _ := TreeFindMax(tree, nil).Content.(Person)
	if p.Age != 18 {
		t.Errorf("found max node error! ", &p)
	}
}

func TestTreeFindMin(t *testing.T) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})

	p, _ := TreeFindMin(tree.Right, nil).Content.(Person)
	if p.Age != 11 {
		t.Errorf("found max node error! ", &p)
	}
}

func TestTreeNodeNum(t *testing.T) {
	tree := TreeRootNew()

	if TreeNodeNum(tree) != 0 {
		t.Error("tree node num is not 0")
	}
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	if TreeNodeNum(tree) != 1 {
		t.Error("tree node num is not 1")
	}
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})
	if TreeNodeNum(tree) != 6 {
		t.Error("tree node num is not 6")
	}
}

func TestTreeLeafNum(t *testing.T) {
	tree := TreeRootNew()

	if TreeLeafNum(tree) != 0 {
		t.Error("tree leaf num is not 0")
	}
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	if TreeLeafNum(tree) != 1 {
		t.Error("tree leaf num is not 1")
	}
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	if TreeLeafNum(tree) != 1 {
		t.Error("tree leaf num is not 1")
	}
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	if TreeLeafNum(tree) != 2 {
		t.Error("tree leaf num is not 2")
	}
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	if TreeLeafNum(tree) != 2 {
		t.Error("tree leaf num is not 2")
	}
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})
	if TreeLeafNum(tree) != 3 {
		t.Error("tree leaf num is not 3")
	}
}

func TestTreeWidth(t *testing.T) {
	tree := TreeRootNew()

	if TreeWidth(tree) != 0 {
		t.Error("tree width num is not 0")
	}
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	if TreeWidth(tree) != 1 {
		t.Error("tree width num is not 1")
	}
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	if TreeWidth(tree) != 1 {
		t.Error("tree width num is not 1")
	}
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	if TreeWidth(tree) != 2 {
		t.Error("tree width num is not 2")
	}
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	if TreeWidth(tree) != 2 {
		t.Error("tree width num is not 2")
	}
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})
	if TreeWidth(tree) != 2 {
		t.Error("tree width num is not 2")
	}
}

func TestTreeLayerNodeNum(t *testing.T) {
	tree := TreeRootNew()

	if TreeLayerNodeNum(tree, 1) != 0 {
		t.Error("tree layer node num is not 0")
	}
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	if TreeLayerNodeNum(tree, 1) != 1 {
		t.Error("tree layer node num is not 1")
	}
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	if TreeLayerNodeNum(tree, 2) != 1 {
		t.Error("tree layer node num is not 1")
	}
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	if TreeLayerNodeNum(tree, 2) != 2 {
		t.Error("tree layer node num is not 2")
	}
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	if TreeLayerNodeNum(tree, 3) != 1 {
		t.Error("tree layer node num is not 1")
	}
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})
	if TreeLayerNodeNum(tree, 4) != 1 {
		t.Error("tree layer node num is not 1")
	}
}

func BenchmarkTreeAdd(b *testing.B) {
	tree := TreeRootNew()
	for i := 0; i < b.N; i++ {
		p := Person{Age: rand.Intn(10000), Name: strconv.Itoa(rand.Intn(10000)) + "Test"}
		if TreeFind(tree, p) == nil {
			TreeAdd(tree, p)
		}
	}
}

func BenchmarkTreeLeafNum(b *testing.B) {
	tree := TreeRootNew()
	TreeAdd(tree, Person{Age: 10, Name: "ZhangSan"})
	TreeAdd(tree, Person{Age: 16, Name: "LiSi"})
	TreeAdd(tree, Person{Age: 6, Name: "WangWu3"})
	TreeAdd(tree, Person{Age: 12, Name: "WangWu"})
	TreeAdd(tree, Person{Age: 18, Name: "WangWu2"})
	TreeAdd(tree, Person{Age: 11, Name: "WangWu4"})

	for i := 0; i < b.N; i++ {
		if TreeLayerNodeNum(tree, 4) != 1 {
			b.Error("tree layer node num is not 1")
		}
	}
}
