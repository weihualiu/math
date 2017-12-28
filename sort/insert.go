package sort

// 插入排序
// 每步将一个待排序的记录，按其顺序码大小插入到前面已排序的字序列的合适位置（从后向前找到合适的位置），直到全部插入排序完为止。

type Insert struct {
	context []Contexter
}

func NewInsert(context ...Contexter) Sorter {
	tmp := new(Insert)
	for _, v := range context {
		if v == nil {
			continue
		}
		tmp.context = append(tmp.context, v)
	}
	return tmp
}

func (this *Insert) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v == nil {
			continue
		}
		this.context = append(this.context, v)
	}
	return this
}

func (this *Insert) Sort() []Contexter {
	length := len(this.context)
	for i := 0; i < length; i++ {
		temp := this.context[i]
		j := i
		for ; j > 0 && !temp.Compare(this.context[j-1]); j-- {
			this.context[j] = this.context[j-1]
		}
		this.context[j] = temp
	}
	return this.context
}
