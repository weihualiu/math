package sort

// 选择排序
//  在要排序的一组值中，选出最小的一个值与第一个位置进行交换；依次在第二个位置之后找出最小的值与第二个位置进行交换。
type Select struct {
	context []Contexter
}

func NewSelect(context ...Contexter) Sorter {
	tmp := new(Select)
	for _, v := range context {
		if v != nil {
			tmp.context = append(tmp.context, v)
		}
	}
	return tmp
}

func (this *Select) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v != nil {
			this.context = append(this.context, v)
		}
	}
	return this
}

func (this *Select) Sort() []Contexter {
	for i := 0; i < len(this.context); i++ {
		//find min value
		var context Contexter
		// 标记出最小值的位置
		var position int
		for j := i; j < len(this.context); j++ {
			if context == nil {
				context = this.context[j]
			}
			if !this.context[j].Compare(context) {
				context = this.context[j]
				position = j
			}
		}
		if i == position {
			continue
		}
		// 用当前值与最小值交换
		this.context[i], this.context[position] = this.context[position], this.context[i]
	}
	return this.context
}
