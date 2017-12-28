package sort

// 冒泡排序
// 重复走访要排序的数列，一次比较两个元素，如果顺序错误则交换这两个元素

type Bubble struct {
	context []Contexter
}

func NewBubble(context ...Contexter) Sorter {
	bubble := new(Bubble)
	for _, v := range context {
		if v != nil {
			bubble.context = append(bubble.context, v)
		}
	}
	return bubble
}

func (this *Bubble) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v != nil {
			this.context = append(this.context, v)
		}
	}
	return this
}

func (this *Bubble) Sort() []Contexter {
	length := len(this.context)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if this.context[j].Compare(this.context[j+1]) {
				this.context[j], this.context[j+1] = this.context[j+1], this.context[j]
			}
		}
	}
	return this.context
}
